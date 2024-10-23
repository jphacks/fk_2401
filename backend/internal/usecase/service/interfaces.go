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
	}

	HouseRepositoryInterface interface {
		CreateHouse(name string) (int64, error)
		GetAllHouses() ([]*domain.House, error)
	}

	ClimateDataRepositoryInterface interface {
		GetAllClimateData() ([]*domain.ClimateData, error)
	}

	UecsDeviceRepositoryInterface interface {
		CreateUecsDevice(newUecsDevice domain.UecsDevice) (int64, error)
		GetAllUecsDevice() ([]*domain.UecsDevice, error)
		GetUecsDeviceFromID(ID int) (*domain.UecsDevice, error)
	}

	M304RepositoryInterface interface {
		CreateM304(newM304 domain.M304) (int64, error)
		GetM304FromUecsDevice(uecsDeviceID int) ([]*domain.M304, error)
	}
)
