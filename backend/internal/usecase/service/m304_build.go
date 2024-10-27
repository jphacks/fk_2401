package service

import (
	"fmt"
	"net/http"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	m304 "github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/M304"
)

type M304BuildService struct {
	deviceRepository     DeviceRepositoryInterface
	uecsDeviceRepository UecsDeviceRepositoryInterface
	m304Repository       M304RepositoryInterface
	m304RecordRepository M304RecordRepositoryInterface
}

func NewM304BuildService(dr DeviceRepositoryInterface, udr UecsDeviceRepositoryInterface, mr M304RepositoryInterface, mrr M304RecordRepositoryInterface) *M304BuildService {
	return &M304BuildService{
		deviceRepository:     dr,
		uecsDeviceRepository: udr,
		m304Repository:       mr,
		m304RecordRepository: mrr,
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

func SettingRly(uecsDeviceID int, m304 domain.M304) *Rly {
	rly := Rly{
		rly_l: 0b00000000,
		rly_h: 0b00000000,
	}
	rly.rly_h = 0b10101010 //allbreak
	rly.rly_l = 0b10101010 //allbreak
	if m304.Rly0 == &uecsDeviceID {
		rly.rly_l += 0b01000000
	}
	if m304.Rly1 == &uecsDeviceID {
		rly.rly_l += 0b00010000
	}
	if m304.Rly2 == &uecsDeviceID {
		rly.rly_l += 0b00000100
	}
	if m304.Rly3 == &uecsDeviceID {
		rly.rly_l += 0b00000001
	}
	if m304.Rly4 == &uecsDeviceID {
		rly.rly_h += 0b01000000
	}
	if m304.Rly5 == &uecsDeviceID {
		rly.rly_h += 0b00010000
	}
	if m304.Rly6 == &uecsDeviceID {
		rly.rly_h += 0b00000100
	}
	if m304.Rly7 == &uecsDeviceID {
		rly.rly_h += 0b00000001
	}
	return &rly
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
	uecsDevice, err := mbs.uecsDeviceRepository.GetUecsDeviceFromID(device.UecsDeviceID)
	if err != nil {
		return 0, err
	}
	m304s, err := mbs.m304Repository.GetM304FromUecsDevice(uecsDevice.ID)
	if err != nil {
		return 0, err
	}
	for _, v := range m304s {
		// データ成型
		valid := 0
		if *device.Valid {
			valid = 1
		} else {
			valid = 255
		}
		ccmtype := uecsDevice.Ccmtype
		uecsData := SettingUECSData(ccmtype)
		rly := SettingRly(uecsDevice.ID, *v)

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
		blockBData := m304.BlockB{
			B_ID:        positionB,
			IP_ADDR:     *v.IpAddr,
			LC_VALID:    valid,
			LC_ROOM:     uecsDevice.Room,
			LC_REGION:   uecsDevice.Region,
			LC_ORDER:    uecsDevice.Order,
			LC_PRIORITY: uecsDevice.Priority,
			LC_LV:       uecsData.lv,
			LC_CAST:     uecsData.cast,
			LC_SR:       reception,
			LC_CCMTYPE:  ccmtype,
			LC_UNIT:     uecsData.unit,
			LC_STHR:     sthr,
			LC_STMN:     stmn,
			LC_EDHR:     edhr,
			LC_EDMN:     edmn,
			LC_INMN:     inmn,
			LC_DUMN:     *device.Duration,
			LC_RLY_L:    rly.rly_l,
			LC_RLY_H:    rly.rly_h,
		}
		blockCData := m304.BlockC{
			C_ID:        positionC,
			IP_ADDR:     *v.IpAddr,
			LC_VALID:    valid,
			LC_ROOM:     uecsDevice.Room,
			LC_REGION:   uecsDevice.Region,
			LC_ORDER:    uecsDevice.Order,
			LC_PRIORITY: uecsDevice.Priority,
			LC_LV:       uecsData.lv,
			LC_CAST:     uecsData.cast,
			LC_SR:       send,
			LC_CCMTYPE:  ccmtype,
			LC_UNIT:     uecsData.unit,
			LC_STHR:     sthr,
			LC_STMN:     stmn,
			LC_EDHR:     edhr,
			LC_EDMN:     edmn,
			LC_INMN:     inmn,
			LC_DUMN:     *device.Duration,
			LC_RLY_L:    rly.rly_l,
			LC_RLY_H:    rly.rly_h,
		}
		blockDData := m304.BlockD{
			D_ID:             positionD,
			IP_ADDR:          *v.IpAddr,
			LC_COPE_VALID:    valid,
			LC_COPE_ROOM:     uecsDevice.Room,
			LC_COPE_REGION:   uecsDevice.Region,
			LC_COPE_ORDER:    uecsDevice.Order,
			LC_COPE_PRIORITY: uecsDevice.Priority,
			LC_COPE_CCMTYPE:  ccmtype,
			LC_COPE_OPE:      *device.Operator,
			LC_COPE_FVAL:     float32(*device.SetPoint),
		}
		// send_param呼び出し
		respB, err := m304.SendBlockB(blockBData)
		if err != nil {
			return 0, err
		}
		for _, w := range respB {
			if w.StatusCode != http.StatusOK {
				return 0, fmt.Errorf("error in response with status code: %d", w.StatusCode)
			}
		}
		respC, err := m304.SendBlockC(blockCData)
		if err != nil {
			return 0, err
		}
		for _, w := range respC {
			if w.StatusCode != http.StatusOK {
				return 0, fmt.Errorf("error in response with status code: %d", w.StatusCode)
			}
		}
		respD, err := m304.SendBlockD(blockDData)
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
