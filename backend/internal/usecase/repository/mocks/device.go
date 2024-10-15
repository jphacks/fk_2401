package mocks

import (
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/repository"
)

type MockDeviceRepository struct {
	DeviceTable       map[int][]*domain.Device
	JoinedDeviceTable map[int][]*repository.JoinedDevice
}

func NewMockDeviceRepository() *MockDeviceRepository {
	deviceTableMap := make(map[int][]*domain.Device)

	device1 := &domain.Device{
		ID:            1,
		HouseID:       1,
		ClimateDataID: 1,
		DeviceName:    "加温器",
		SetPoint:      22,
		Duration:      1,
	}
	device2 := &domain.Device{
		ID:            2,
		HouseID:       2,
		ClimateDataID: 2,
		DeviceName:    "加湿器",
		SetPoint:      60,
		Duration:      2,
	}
	device3 := &domain.Device{
		ID:            3,
		HouseID:       2,
		ClimateDataID: 3,
		DeviceName:    "CO2濃度センサ",
		SetPoint:      420,
		Duration:      3,
	}
	device4 := &domain.Device{
		ID:            4,
		HouseID:       3,
		ClimateDataID: 3,
		DeviceName:    "CO2濃度センサ",
		SetPoint:      420,
		Duration:      4,
	}

	deviceTableMap[1] = []*domain.Device{device1}
	deviceTableMap[2] = []*domain.Device{device2, device3}
	deviceTableMap[3] = []*domain.Device{device4}

	joinedDeviceTableMap := make(map[int][]*repository.JoinedDevice)

	joinedDevice1 := &repository.JoinedDevice{
		ID:          1,
		HouseID:     1,
		SetPoint:    22,
		Duration:    1,
		ClimateData: "気温",
		Unit:        "℃",
	}
	joinedDevice2 := &repository.JoinedDevice{
		ID:          2,
		HouseID:     2,
		SetPoint:    60,
		Duration:    2,
		ClimateData: "湿度",
		Unit:        "%",
	}
	joinedDevice3 := &repository.JoinedDevice{
		ID:          3,
		HouseID:     2,
		SetPoint:    420,
		Duration:    3,
		ClimateData: "二酸化炭素量",
		Unit:        "ppm",
	}
	joinedDevice4 := &repository.JoinedDevice{
		ID:          4,
		HouseID:     3,
		SetPoint:    420,
		Duration:    4,
		ClimateData: "二酸化炭素量",
		Unit:        "ppm",
	}

	joinedDeviceTableMap[1] = []*repository.JoinedDevice{joinedDevice1}
	joinedDeviceTableMap[2] = []*repository.JoinedDevice{joinedDevice2, joinedDevice3}
	joinedDeviceTableMap[3] = []*repository.JoinedDevice{joinedDevice4}

	return &MockDeviceRepository{
		DeviceTable:       deviceTableMap,
		JoinedDeviceTable: joinedDeviceTableMap,
	}

}

func (dr MockDeviceRepository) GetDevicesFromHouse(houseID int) ([]*domain.Device, error) {
	return dr.DeviceTable[houseID], nil
}

func (dr MockDeviceRepository) GetJoinedDevicesFromHouse(houseID int) ([]*repository.JoinedDevice, error) {
	return dr.JoinedDeviceTable[houseID], nil
}
