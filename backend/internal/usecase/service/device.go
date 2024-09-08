package service

import (
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
)

type DeviceService struct {
	deviceRepository DeviceRepositoryInterface
}

func (ds DeviceService) GetDevicesFromHouse(houseID int) ([]*domain.Device, error) {
	devices, err := ds.deviceRepository.GetDevicesFromHouse(houseID)
	if err != nil {
		return nil, err
	}

	return devices, err
}
