package service

import (
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/repository"
)

type (
	// Repository interfaces
	DeviceRepositoryInterface interface {
		GetDevicesFromHouse(houseID int) ([]*domain.Device, error)
		GetJoinedDevicesFromHouse(houseID int) ([]*repository.JoinedDevice, error)
	}

	HouseRepositoryInterface interface {
		GetAllHouses() ([]*domain.House, error)
	}

	// ClimateDataRepositoryInterface interface {
	// }
)
