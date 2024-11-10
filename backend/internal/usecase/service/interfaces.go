package service

import (
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/repository"
)

type (
	// Repository interfaces
	DeviceRepositoryInterface interface {
		CreateDevice(newDevice domain.Device) (int64, error)
		GetDevicesFromHouse(houseID int) ([]*domain.Device, error)
		GetJoinedDevicesFromHouse(houseID int) ([]*repository.JoinedDevice, error)
		GetDeviceFromID(deviceID int) (*domain.Device, error)
	}

	HouseRepositoryInterface interface {
		CreateHouse(name string) (int64, error)
		GetAllHouses() ([]*domain.House, error)
	}

	ClimateDataRepositoryInterface interface {
		GetAllClimateData() ([]*domain.ClimateData, error)
		GetClimateDataFromID(ID int) (*domain.ClimateData, error)
	}

	SensorRepositoryInterface interface {
		CreateSensor(newSensor domain.Sensor) (int64, error)
		GetAllSensor() ([]*domain.Sensor, error)
		GetSensorFromID(ID int) (*domain.Sensor, error)
	}

	M304RepositoryInterface interface {
		CreateM304(newM304 domain.M304) (int64, error)
		GetM304FromID(ID int) ([]*domain.M304, error)
	}

	M304RecordRepositoryInterface interface {
		CreateM304Record(NewM304Record domain.M304Record) (int64, error)
		GetM304RecordFromM304ID(m304ID int) ([]*domain.M304Record, error)
	}

	DeviceConditionRepositoryInterface interface {
		CreateDeviceCondition(newDeviceCondition domain.DeviceCondition) (int64, error)
		GetDeviceConditionFromID(ID int) (*domain.DeviceCondition, error)
		GetDeviceConditionsFromDeviceID(deviceID int) ([]*domain.DeviceCondition, error)
	}

	OperationRepositoryInterface interface {
		CreateOperation(newOperation domain.Operation) (int64, error)
		GetOperationFromID(ID int) (*domain.Operation, error)
		GetOperationsFromDeviceID(deviceID int) ([]*domain.Operation, error)
	}

	TimeScheduleRepositoryInterface interface {
		CreateTimeSchedule(newTimeSchedule domain.TimeSchedule) (int64, error)
		GetTimeScheduleFromID(ID int) (*domain.TimeSchedule, error)
		GetTimeSchedulesFromDeviceCondition(deviceConditionID int) ([]*domain.TimeSchedule, error)
	}
)
