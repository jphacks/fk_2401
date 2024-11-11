package service

import (
	"fmt"
	"math"
	"net/http"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	m304_send "github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/M304"
)

type M304BuildService struct {
	deviceRepository          DeviceRepositoryInterface
	sensorRepository          SensorRepositoryInterface
	m304Repository            M304RepositoryInterface
	deviceConditionRepository DeviceConditionRepositoryInterface
	operationRepository       OperationRepositoryInterface
	timeScheduleRepository    TimeScheduleRepositoryInterface
	m304RecordRepository      M304RecordRepositoryInterface
}

func NewM304BuildService(dr DeviceRepositoryInterface, sr SensorRepositoryInterface, mr M304RepositoryInterface, dcr DeviceConditionRepositoryInterface, or OperationRepositoryInterface, tsr TimeScheduleRepositoryInterface, mrr M304RecordRepositoryInterface) *M304BuildService {
	return &M304BuildService{
		deviceRepository:          dr,
		sensorRepository:          sr,
		m304Repository:            mr,
		deviceConditionRepository: dcr,
		operationRepository:       or,
		timeScheduleRepository:    tsr,
		m304RecordRepository:      mrr,
	}
}

type UECSData struct {
	unit string
	cast int //小数点位置
	lv   int
}

const (
	Brecnum   = 30
	Crecnum   = 10
	Drecnum   = 10
	A1S0      = 1
	A1S1      = 2
	A10S0     = 3
	A10S1     = 4
	A1M0      = 5
	A1M1      = 6
	B0        = 7
	B1        = 8
	S1S0      = 9
	S1M0      = 10
	sthr      = 0
	stmn      = 0
	edhr      = 23
	edmn      = 59
	inmn      = 0
	send      = "S"
	reception = "R"
)

func SettingUECSData(ccmtype string) *UECSData {
	uecsData := UECSData{
		unit: "",
		cast: 0,
		lv:   A1M0,
	}
	switch ccmtype {
	case "Time":
		uecsData.unit = "hms"
		uecsData.cast = 0
		uecsData.lv = A1M0
	case "Date":
		uecsData.unit = "ymd"
		uecsData.cast = 0
		uecsData.lv = A1M0
	case "InAirTemp":
		uecsData.unit = "C"
		uecsData.cast = 1
		uecsData.lv = A10S0
	case "InAirHumid":
		uecsData.unit = "%"
		uecsData.cast = 0
		uecsData.lv = A10S0
	case "InAirCO2":
		uecsData.unit = "ppm"
		uecsData.cast = 0
		uecsData.lv = A10S0
	case "InRadiation":
		uecsData.unit = "kW m-2"
		uecsData.cast = 2
		uecsData.lv = A10S0
	case "WAirTemp":
		uecsData.unit = "C"
		uecsData.cast = 1
		uecsData.lv = A10S0
	case "WAirHumid":
		uecsData.unit = "%"
		uecsData.cast = 0
		uecsData.lv = A10S0
	case "WAirCO2":
		uecsData.unit = "ppm"
		uecsData.cast = 0
		uecsData.lv = A10S0
	case "WRadiation":
		uecsData.unit = "kW m-2"
		uecsData.cast = 2
		uecsData.lv = A10S0
	case "WWindSpeed":
		uecsData.unit = "m s-1"
		uecsData.cast = 0
		uecsData.lv = A10S0
	case "WWindDir16":
		uecsData.unit = ""
		uecsData.cast = 0
		uecsData.lv = A10S0
	case "WRainfall":
		uecsData.unit = ""
		uecsData.cast = 0
		uecsData.lv = A10S0
	default:
		uecsData.unit = ""
		uecsData.cast = 0
		uecsData.lv = A1M0
	}
	return &uecsData
}

type Rly struct {
	rly_l int
	rly_h int
}

// rly 0~7, rly_on 00 ~ 11 -> 0 ~ 3
func SettingRly(rly int, rly_on int) *Rly {
	rly_struct := Rly{
		rly_l: 0b00000000,
		rly_h: 0b00000000,
	}
	if rly < 4 {
		rly_struct.rly_l += int(math.Pow(4.0, (3.0-float64(rly))) * float64(rly_on))
	} else if rly < 8 {
		rly_struct.rly_h += int(math.Pow(4.0, (7.0-float64(rly))) * float64(rly_on))
	}
	return &rly_struct
}

func findMissingNumbers(arr []int, num int) []int {
	fullSet := make(map[int]bool)
	for i := 0; i < num; i++ {
		fullSet[i] = true
	}

	for _, num := range arr {
		delete(fullSet, num)
	}

	missingNumbers := []int{}
	for num := range fullSet {
		missingNumbers = append(missingNumbers, num)
	}
	return missingNumbers
}

func GetBuildAddress(rec_num int, arr []*domain.M304Record) (int, error) {
	count := 0
	var val_arr []int
	for _, v := range arr {
		if v.Valid {
			count += 1
			val_arr = append(val_arr, v.Position)
		}
	}
	if count >= rec_num {
		return rec_num, fmt.Errorf("no room in the record")
	} else {
		missingNumbers := findMissingNumbers(val_arr, rec_num)
		if len(missingNumbers) == 0 {
			return 0, fmt.Errorf("no missing number found")
		}
		position := missingNumbers[0]
		return position, nil
	}
}

func (mbs M304BuildService) BuildM304(deviceID int) (int, error) {
	device, err := mbs.deviceRepository.GetDeviceFromID(deviceID)
	if err != nil {
		return 0, err
	}
	sensor, err := mbs.sensorRepository.GetSensorFromID(device.SensorID)
	if err != nil {
		return 0, err
	}
	m304, err := mbs.m304Repository.GetM304FromID(device.M304ID)
	if err != nil {
		return 0, err
	}
	deviceConditions, err := mbs.deviceConditionRepository.GetDeviceConditionsFromDeviceID(deviceID)
	if err != nil {
		return 0, err
	}
	operations, err := mbs.operationRepository.GetOperationsFromDeviceID(deviceID)
	if err != nil {
		return 0, err
	}
	timeSchedules := make([][]*domain.TimeSchedule, len(operations))
	for i, v := range operations {
		getTimeSchedules, err := mbs.timeScheduleRepository.GetTimeSchedulesFromDeviceCondition(v.ID)
		if err != nil {
			return 0, err
		}
		timeSchedules[i] = getTimeSchedules
	}

	for i, v := range deviceConditions {
		// データ成型
		valid := 0
		if v.Valid {
			valid = 1
		} else {
			valid = 255
		}
		ccmtype := sensor.Ccmtype
		uecsData := SettingUECSData(ccmtype)
		// 要変更
		rly := SettingRly(*device.Rly, operations)

		// 書き込み位置確認
		m304records, err := mbs.m304RecordRepository.GetM304RecordFromM304ID(v.ID)
		if err != nil {
			return 0, err
		}
		var B []*domain.M304Record
		var C []*domain.M304Record
		var D []*domain.M304Record
		for _, w := range m304records {
			if w.Block == "B" {
				B = append(B, w)
			} else if w.Block == "C" {
				C = append(C, w)
			} else if w.Block == "D" {
				D = append(D, w)
			}
		}
		positionB, err := GetBuildAddress(Brecnum, B)
		if err != nil {
			return 0, err
		}
		positionC, err := GetBuildAddress(Crecnum, C)
		if err != nil {
			return 0, err
		}
		positionD, err := GetBuildAddress(Drecnum, D)
		if err != nil {
			return 0, err
		}
		blockBData := m304_send.BlockB{
			BID:        positionB,
			IpAddr:     *m304.IpAddr,
			LcValid:    valid,
			LcRoom:     sensor.Room,
			LcRegion:   sensor.Region,
			LcOrder:    sensor.Order,
			LcPriority: sensor.Priority,
			LcLv:       uecsData.lv,
			LcCast:     uecsData.cast,
			LcSr:       reception,
			LcCcmType:  ccmtype,
			LcUnit:     uecsData.unit,
			LcStHr:     sthr,
			LcStMn:     stmn,
			LcEdHr:     edhr,
			LcEdMn:     edmn,
			LcInMn:     inmn,
			LcDuMn:     *v.Duration,
			LcRlyL:     rly.rly_l,
			LcRlyH:     rly.rly_h,
		}
		blockCData := m304_send.BlockC{
			CID:        positionC,
			IpAddr:     *m304.IpAddr,
			LcValid:    valid,
			LcRoom:     sensor.Room,
			LcRegion:   sensor.Region,
			LcOrder:    sensor.Order,
			LcPriority: sensor.Priority,
			LcLv:       uecsData.lv,
			LcCast:     uecsData.cast,
			LcSr:       send,
			LcCcmType:  ccmtype,
			LcUnit:     uecsData.unit,
			LcStHr:     sthr,
			LcStMn:     stmn,
			LcEdHr:     edhr,
			LcEdMn:     edmn,
			LcInMn:     inmn,
			LcDuMn:     *v.Duration,
			LcRlyL:     rly.rly_l,
			LcRlyH:     rly.rly_h,
		}
		blockDData := m304_send.BlockD{
			DID:            positionD,
			IpAddr:         *m304.IpAddr,
			LcCopeValid:    valid,
			LcCopeRoom:     sensor.Room,
			LcCopeRegion:   sensor.Region,
			LcCopeOrder:    sensor.Order,
			LcCopePriority: sensor.Priority,
			LcCopeCcmType:  ccmtype,
			LcCopeOpe:      *v.Operator,
			LcCopeFval:     float32(*v.SetPoint),
		}
		// send_param呼び出し
		respB, err := m304_send.SendBlockB(blockBData)
		if err != nil {
			return 0, err
		}
		for _, w := range respB {
			if w.StatusCode != http.StatusOK {
				return 0, fmt.Errorf("error in response with status code: %d", w.StatusCode)
			}
		}
		respC, err := m304_send.SendBlockC(blockCData)
		if err != nil {
			return 0, err
		}
		for _, w := range respC {
			if w.StatusCode != http.StatusOK {
				return 0, fmt.Errorf("error in response with status code: %d", w.StatusCode)
			}
		}
		respD, err := m304_send.SendBlockD(blockDData)
		if err != nil {
			return 0, err
		}
		for _, w := range respD {
			if w.StatusCode != http.StatusOK {
				return 0, fmt.Errorf("error in response with status code: %d", w.StatusCode)
			}
		}
	}
	return 1, nil
}
