package repository

import (
	"context"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
)

type SensorRepository struct {
	queries *mysqlc.Queries
}

func NewSensorRepository(queries *mysqlc.Queries) *SensorRepository {
	return &SensorRepository{
		queries: queries,
	}
}

func (udr SensorRepository) CreateSensor(newSensor domain.Sensor) (int64, error) {
	ctx := context.Background()

	arg := mysqlc.CreateSensorParams{
		CcmType:  newSensor.Ccmtype,
		Room:     int32(newSensor.Room),
		Region:   int32(newSensor.Region),
		Order:    int32(newSensor.Order),
		Priority: int32(newSensor.Priority),
	}

	id, err := udr.queries.CreateSensor(ctx, arg)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (udr SensorRepository) GetAllSensor() ([]*domain.Sensor, error) {
	ctx := context.Background()

	SensorRow, err := udr.queries.GetAllSensor(ctx)
	if err != nil {
		return nil, err
	}

	Sensors := make([]*domain.Sensor, len(SensorRow))
	for i, v := range SensorRow {
		Sensors[i] = domain.NewSensorWithID(
			int(v.ID),
			v.CcmType,
			int(v.Room),
			int(v.Region),
			int(v.Order),
			int(v.Priority),
		)
	}

	return Sensors, nil
}

func (udr SensorRepository) GetSensorFromID(ID int) (*domain.Sensor, error) {
	ctx := context.Background()

	getSensor, err := udr.queries.GetSensorFromID(ctx, int32(ID))
	if err != nil {
		return nil, err
	}

	Sensor := domain.NewSensorWithID(
		int(getSensor.ID),
		getSensor.CcmType,
		int(getSensor.Room),
		int(getSensor.Region),
		int(getSensor.Order),
		int(getSensor.Priority),
	)

	return Sensor, nil
}
