package repository

import (
	"context"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
)

type DeviceRepository struct {
	queries *mysqlc.Queries
}

func (dr DeviceRepository) GetDevicesFromHouse(houseID int) ([]*domain.Device, error) {
	ctx := context.Background()

	devicesRow, err := dr.queries.GetDevicesFromHouse(ctx, int32(houseID))
	if err != nil {
		return nil, err
	}

	devices := make([]*domain.Device, len(devicesRow))
	for i, v := range devicesRow {
		devices[i] = domain.NewDeviceWithID(int(v.HouseID))
	}
}
