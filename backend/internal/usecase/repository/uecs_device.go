package repository

import (
	"context"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
)

type UecsDeviceRepository struct {
	queries *mysqlc.Queries
}

func NewUecsDeviceRepository(queries *mysqlc.Queries) *UecsDeviceRepository {
	return &UecsDeviceRepository{
		queries: queries,
	}
}

func (udr UecsDeviceRepository) CreateUecsDevice(newUecsDevice domain.UecsDevice) (int64, error) {
	ctx := context.Background()

	arg := mysqlc.CreateUecsDeviceParams{
		Ccmtype:  newUecsDevice.Ccmtype,
		Room:     int32(newUecsDevice.Room),
		Region:   int32(newUecsDevice.Region),
		Order:    int32(newUecsDevice.Order),
		Priority: int32(newUecsDevice.Priority),
	}

	id, err := udr.queries.CreateUecsDevice(ctx, arg)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (udr UecsDeviceRepository) GetAllUecsDevice() ([]*domain.UecsDevice, error) {
	ctx := context.Background()

	uecsDeviceRow, err := udr.queries.GetAllUecsDevice(ctx)
	if err != nil {
		return nil, err
	}

	uecsDevices := make([]*domain.UecsDevice, len(uecsDeviceRow))
	for i, v := range uecsDeviceRow {
		uecsDevices[i] = domain.NewUecsDeviceWithID(
			int(v.ID),
			v.Ccmtype,
			int(v.Room),
			int(v.Region),
			int(v.Order),
			int(v.Priority),
		)
	}

	return uecsDevices, nil
}

func (udr UecsDeviceRepository) GetUecsDeviceFromID(ID int) (*domain.UecsDevice, error) {
	ctx := context.Background()

	getUecsDevice, err := udr.queries.GetUecsDeviceFromID(ctx, int32(ID))
	if err != nil {
		return nil, err
	}

	uecsDevice := domain.NewUecsDeviceWithID(
		int(getUecsDevice.ID),
		getUecsDevice.Ccmtype,
		int(getUecsDevice.Room),
		int(getUecsDevice.Region),
		int(getUecsDevice.Order),
		int(getUecsDevice.Priority),
	)

	return uecsDevice, nil
}
