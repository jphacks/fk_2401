package service

import (
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/repository"
)

type DeviceService struct {
	deviceRepository DeviceRepositoryInterface
}

func NewDeviceService(dr *repository.DeviceRepository) *DeviceService {
	return &DeviceService{
		deviceRepository: dr,
	}
}

// デバイスのみを取得するメソッド
func (ds DeviceService) GetDevices(houseID int) ([]*domain.Device, error) {
	devices, err := ds.deviceRepository.GetDevicesFromHouse(houseID)
	if err != nil {
		return nil, err
	}

	return devices, nil
}

// デバイスと気象データを1セットとして取得するメソッド
func (ds DeviceService) GetJoinedDevices(houseID int) ([]*repository.JoinedDevice, error) {
	joinedDevices, err := ds.deviceRepository.GetJoinedDevicesFromHouse(houseID)
	if err != nil {
		return nil, err
	}

	return joinedDevices, nil
}
