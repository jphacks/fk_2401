package service

import (
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/repository"
)

type DeviceService struct {
	deviceRepository DeviceRepositoryInterface
}

func NewDeviceService(dr DeviceRepositoryInterface) *DeviceService {
	return &DeviceService{
		deviceRepository: dr,
	}
}

func (ds DeviceService) CreateDevice(newDevice domain.Device) (int64, error) {
	id, err := ds.deviceRepository.CreateDevice(newDevice)
	if err != nil {
		return 0, err
	}

	return id, nil
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

func (ds DeviceService) GetDeviceFromID(ID int) (*domain.Device, error) {
	device, err := ds.deviceRepository.GetDeviceFromID(ID)
	if err != nil {
		return nil, err
	}

	return device, nil
}
