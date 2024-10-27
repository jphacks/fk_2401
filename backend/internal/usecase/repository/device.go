package repository

import (
	"context"
	"database/sql"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
)

type JoinedDevice struct {
	ID           int
	HouseID      int
	UecsDeviceID int
	DeviceName   *string
	Valid        *bool
	SetPoint     *float64
	Duration     *int
	Operator     *int
	ClimateData  string
	Unit         string
}

func NewJoinedDevice(id int, houseID int, uecsDeviceID int, deviceName *string, valid *bool, setPoint *float64, duration *int, operator *int, climateData, unit string) *JoinedDevice {
	return &JoinedDevice{
		ID:           id,
		HouseID:      houseID,
		UecsDeviceID: uecsDeviceID,
		DeviceName:   deviceName,
		Valid:        valid,
		SetPoint:     setPoint,
		Duration:     duration,
		Operator:     operator,
		ClimateData:  climateData,
		Unit:         unit,
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
		UecsDeviceID:  int32(newDevice.UecsDeviceID),
		DeviceName: sql.NullString{
			String: func() string {
				if newDevice.DeviceName != nil {
					return *newDevice.DeviceName
				}
				return ""
			}(),
			Valid: newDevice.DeviceName != nil,
		},
		Valid: sql.NullBool{
			Bool: func() bool {
				if newDevice.Valid != nil {
					return *newDevice.Valid
				}
				return false
			}(),
			Valid: newDevice.Valid != nil,
		},
		SetPoint: sql.NullFloat64{
			Float64: func() float64 {
				if newDevice.SetPoint != nil {
					return *newDevice.SetPoint
				}
				return 0
			}(),
			Valid: newDevice.SetPoint != nil,
		},
		Duration: sql.NullInt32{
			Int32: func() int32 {
				if newDevice.Duration != nil {
					return int32(*newDevice.Duration)
				}
				return 0
			}(),
			Valid: newDevice.Duration != nil,
		},
		Operator: sql.NullInt32{
			Int32: func() int32 {
				if newDevice.Operator != nil {
					return int32(*newDevice.Operator)
				}
				return 0
			}(),
			Valid: newDevice.Operator != nil,
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
		var deviceName *string
		if v.DeviceName.Valid {
			dN := v.DeviceName.String
			deviceName = &dN
		}
		var valid *bool
		if v.Valid.Valid {
			vl := v.Valid.Bool
			valid = &vl
		}
		var setPoint *float64
		if v.SetPoint.Valid {
			sP := v.SetPoint.Float64
			setPoint = &sP
		}
		var duration *int
		if v.Duration.Valid {
			du := int(v.Duration.Int32)
			duration = &du
		}
		var operator *int
		if v.Operator.Valid {
			ope := int(v.Operator.Int32)
			operator = &ope
		}
		devices[i] = domain.NewDeviceWithID(
			int(v.ID),
			int(v.HouseID),
			int(v.ClimateDataID),
			int(v.UecsDeviceID),
			deviceName,
			valid,
			setPoint,
			duration,
			operator,
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
		var deviceName *string
		if v.DeviceName.Valid {
			dN := v.DeviceName.String
			deviceName = &dN
		}
		var valid *bool
		if v.Valid.Valid {
			vl := v.Valid.Bool
			valid = &vl
		}
		var setPoint *float64
		if v.SetPoint.Valid {
			sP := v.SetPoint.Float64
			setPoint = &sP
		}
		var duration *int
		if v.Duration.Valid {
			du := int(v.Duration.Int32)
			duration = &du
		}
		var operator *int
		if v.Operator.Valid {
			ope := int(v.Operator.Int32)
			operator = &ope
		}
		joinedDevices[i] = NewJoinedDevice(
			int(v.ID),
			int(v.HouseID),
			int(v.UecsDeviceID),
			deviceName,
			valid,
			setPoint,
			duration,
			operator,
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

	var deviceName *string
	if device.DeviceName.Valid {
		dN := device.DeviceName.String
		deviceName = &dN
	}
	var valid *bool
	if device.Valid.Valid {
		vl := device.Valid.Bool
		valid = &vl
	}
	var setPoint *float64
	if device.SetPoint.Valid {
		sP := device.SetPoint.Float64
		setPoint = &sP
	}
	var duration *int
	if device.Duration.Valid {
		du := int(device.Duration.Int32)
		duration = &du
	}
	var operator *int
	if device.Operator.Valid {
		ope := int(device.Operator.Int32)
		operator = &ope
	}
	getDevice := domain.NewDeviceWithID(
		int(device.ID),
		int(device.HouseID),
		int(device.ClimateDataID),
		int(device.UecsDeviceID),
		deviceName,
		valid,
		setPoint,
		duration,
		operator,
	)

	return getDevice, nil
}
