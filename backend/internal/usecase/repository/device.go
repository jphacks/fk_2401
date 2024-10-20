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
	DeviceName  string
	SetPoint    float64
	Duration    int
	ClimateData string
	Unit        string
}

func NewJoinedDevice(id int, houseID int, deviceName string, setPoint float64, duration int, climateData, unit string) *JoinedDevice {
	return &JoinedDevice{
		ID:          id,
		HouseID:     houseID,
		DeviceName:  deviceName,
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

	setpoint := sql.NullFloat64{
		Float64: newDevice.SetPoint,
		Valid:   false,
	}
	if newDevice.SetPoint != 0 {
		setpoint.Valid = true
	}
	duration := sql.NullInt32{
		Int32: int32(newDevice.Duration),
		Valid: false,
	}
	if newDevice.Duration != 0 {
		duration.Valid = true
	}

	arg := mysqlc.CreateDeviceParams{
		HouseID:       int32(newDevice.HouseID),
		ClimateDataID: int32(newDevice.ClimateDataID),
		DeviceName: sql.NullString{
			String: newDevice.DeviceName,
			Valid:  true,
		},
		SetPoint: setpoint,
		Duration: duration,
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
		setpoint := float64(0)
		if v.SetPoint.Valid {
			setpoint = v.SetPoint.Float64
		}
		duration := 0
		if v.Duration.Valid {
			duration = int(v.Duration.Int32)
		}
		devices[i] = domain.NewDeviceWithID(
			int(v.ID),
			int(v.HouseID),
			int(v.ClimateDataID),
			v.DeviceName.String,
			setpoint,
			duration,
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
		setpoint := float64(0)
		if v.SetPoint.Valid {
			setpoint = v.SetPoint.Float64
		}
		duration := 0
		if v.Duration.Valid {
			duration = int(v.Duration.Int32)
		}
		joinedDevices[i] = NewJoinedDevice(
			int(v.ID),
			int(v.HouseID),
			v.DeviceName.String,
			setpoint,
			duration,
			v.ClimateDataName,
			v.Unit,
		)
	}

	return joinedDevices, nil
}
