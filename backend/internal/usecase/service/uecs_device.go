package service

import "github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"

type UecsDeviceService struct {
	uecsDeviceRepository UecsDeviceRepositoryInterface
}

func NewUecsDeviceService(udr UecsDeviceRepositoryInterface) *UecsDeviceService {
	return &UecsDeviceService{
		uecsDeviceRepository: udr,
	}
}

func (uds UecsDeviceService) CreateUecsDevice(newUecsDevice domain.UecsDevice) (int64, error) {
	id, err := uds.uecsDeviceRepository.CreateUecsDevice(newUecsDevice)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (uds UecsDeviceService) GetAllUecsDevice() ([]*domain.UecsDevice, error) {
	uecsDevices, err := uds.uecsDeviceRepository.GetAllUecsDevice()
	if err != nil {
		return nil, err
	}

	return uecsDevices, nil
}

func (uds UecsDeviceService) GetUecsDeviceFromID(ID int) (*domain.UecsDevice, error) {
	uecsDevices, err := uds.uecsDeviceRepository.GetUecsDeviceFromID(ID)
	if err != nil {
		return nil, err
	}

	return uecsDevices, nil
}
