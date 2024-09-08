package service

import "github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"

type (
	// Repository interfaces
	DeviceRepositoryInterface interface {
		GetDevicesFromHouse(houseID int) ([]*domain.Device, error)
	}

	// HouseRepositoryInterface interface {
	// }

	// ClimateDataRepositoryInterface interface {
	// }
)
