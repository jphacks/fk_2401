package repository

import (
	"context"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
)

type JoinedDevice struct {
	ID          int
	HouseID     int
	M304ID      int
	SensorID    int
	DeviceName  *string
	Rly         *int
	ClimateData string
	Unit        string
}

func NewJoinedDevice(id int, houseID int, m304ID int, sensorID int, deviceName *string, rly *int, climateData, unit string) *JoinedDevice {
	return &JoinedDevice{
		ID:          id,
		HouseID:     houseID,
		M304ID:      m304ID,
		SensorID:    sensorID,
		DeviceName:  deviceName,
		Rly:         rly,
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
		M304ID:        int32(newDevice.M304ID),
		SensorID:      int32(newDevice.SensorID),
		DeviceName:    PointerToNullString(newDevice.DeviceName),
		Rly:           PointerToNullInt32(newDevice.Rly),
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
			int(v.M304ID),
			int(v.SensorID),
			NullStringToPointer(v.DeviceName),
			NullInt32ToPointer(v.Rly),
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
			int(v.M304ID),
			int(v.SensorID),
			NullStringToPointer(v.DeviceName),
			NullInt32ToPointer(v.Rly),
			v.ClimateDataName,
			v.Unit,
		)
	}

	return joinedDevices, nil
}

func (dr DeviceRepository) GetDeviceFromID(ID int) (*domain.Device, error) {
	ctx := context.Background()

	device, err := dr.queries.GetDeviceFromID(ctx, int32(ID))
	if err != nil {
		return nil, err
	}

	getDevice := domain.NewDeviceWithID(
		int(device.ID),
		int(device.HouseID),
		int(device.ClimateDataID),
		int(device.M304ID),
		int(device.SensorID),
		NullStringToPointer(device.DeviceName),
		NullInt32ToPointer(device.Rly),
	)

	return getDevice, nil
}
