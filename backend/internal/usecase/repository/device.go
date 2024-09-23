package repository

import (
	"context"
	"database/sql"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
)

type JoinedDevice struct {
	ID          int
	HouseID     int
	SetPoint    int
	Duration    int
	ClimateData string
	Unit        string
}

func NewJoinedDevice(id, houseID, setPoint, duration int, climateData, unit string) *JoinedDevice {
	return &JoinedDevice{
		ID:          id,
		HouseID:     houseID,
		SetPoint:    setPoint,
		Duration:    duration,
		ClimateData: climateData,
		Unit:        unit,
	}
}

type DeviceRepository struct {
	queries *mysqlc.Queries
}

func NewDeviceRepository(queries *mysqlc.Queries) *DeviceRepository {
	return &DeviceRepository{
		queries: queries,
	}
}

func (dr DeviceRepository) CreateDevice(newDevice domain.Device) (int64, error) {
	ctx := context.Background()

	arg := mysqlc.CreateDeviceParams{
		HouseID:       int32(newDevice.HouseID),
		ClimateDataID: int32(newDevice.ClimateDataID),
		SetPoint: sql.NullInt32{
			Int32: int32(newDevice.SetPoint),
			Valid: true,
		},
		Duration: sql.NullInt32{
			Int32: int32(newDevice.Duration),
			Valid: true,
		},
	}

	id, err := dr.queries.CreateDevice(ctx, arg)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (dr DeviceRepository) GetDevicesFromHouse(houseID int) ([]*domain.Device, error) {
	ctx := context.Background()

	devicesRow, err := dr.queries.GetDevicesFromHouse(ctx, int32(houseID))
	if err != nil {
		return nil, err
	}

	devices := make([]*domain.Device, len(devicesRow))
	for i, v := range devicesRow {
		devices[i] = domain.NewDeviceWithID(
			int(v.ID),
			int(v.HouseID),
			int(v.ClimateDataID),
			int(v.SetPoint.Int32),
			int(v.Duration.Int32),
		)
	}

	return devices, nil
}

func (dr DeviceRepository) GetJoinedDevicesFromHouse(houseID int) ([]*JoinedDevice, error) {
	ctx := context.Background()

	joinedDevicesRow, err := dr.queries.GetJoinedDevicesFromHouse(ctx, int32(houseID))
	if err != nil {
		return nil, err
	}

	joinedDevices := make([]*JoinedDevice, len(joinedDevicesRow))
	for i, v := range joinedDevicesRow {
		joinedDevices[i] = NewJoinedDevice(
			int(v.ID),
			int(v.HouseID),
			int(v.SetPoint.Int32),
			int(v.Duration.Int32),
			v.ClimateDataName,
			v.Unit,
		)
	}

	return joinedDevices, nil
}
