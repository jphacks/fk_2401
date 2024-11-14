package service

import "github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"

type SensorService struct {
	sensorRepository SensorRepositoryInterface
}

func NewSensorService(sr SensorRepositoryInterface) *SensorService {
	return &SensorService{
		sensorRepository: sr,
	}
}

func (ss SensorService) CreateSensor(newSensor domain.Sensor) (int64, error) {
	id, err := ss.sensorRepository.CreateSensor(newSensor)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ss SensorService) GetAllSensor() ([]*domain.Sensor, error) {
	sensors, err := ss.sensorRepository.GetAllSensor()
	if err != nil {
		return nil, err
	}

	return sensors, nil
}

func (ss SensorService) GetSensorFromID(ID int) (*domain.Sensor, error) {
	sensors, err := ss.sensorRepository.GetSensorFromID(ID)
	if err != nil {
		return nil, err
	}

	return sensors, nil
}
