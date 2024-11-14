package service

import "github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"

type DeviceConditionService struct {
	deviceConditionRepository DeviceConditionRepositoryInterface
}

func NewDeviceConditionService(dcr DeviceConditionRepositoryInterface) *DeviceConditionService {
	return &DeviceConditionService{
		deviceConditionRepository: dcr,
	}
}

func (dcs DeviceConditionService) CreateDeviceCondition(newDeviceCondition domain.DeviceCondition) (int64, error) {
	id, err := dcs.deviceConditionRepository.CreateDeviceCondition(newDeviceCondition)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dcs DeviceConditionService) GetDeviceConditionFromID(ID int) (*domain.DeviceCondition, error) {
	deviceCondition, err := dcs.deviceConditionRepository.GetDeviceConditionFromID(ID)
	if err != nil {
		return nil, err
	}

	return deviceCondition, nil
}

func (dcs DeviceConditionService) GetDeviceConditionsFromDeviceID(deviceID int) ([]*domain.DeviceCondition, error) {
	deviceConditions, err := dcs.deviceConditionRepository.GetDeviceConditionsFromDeviceID(deviceID)
	if err != nil {
		return nil, err
	}

	return deviceConditions, nil
}
