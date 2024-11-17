package repository

import (
	"context"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
)

type DeviceConditionRepository struct {
	queries *mysqlc.Queries
}

func NewDeviceConditionRepository(queries *mysqlc.Queries) *DeviceConditionRepository {
	return &DeviceConditionRepository{
		queries: queries,
	}
}

func (dcr DeviceConditionRepository) CreateDeviceCondition(newDeviceCondition domain.DeviceCondition) (int64, error) {
	ctx := context.Background()

	arg := mysqlc.CreateDeviceConditionParams{
		DeviceID:    int32(newDeviceCondition.DeviceID),
		OperationID: int32(newDeviceCondition.OperationID),
		Valid:       newDeviceCondition.Valid,
		SetPoint:    PointerToNullFloat64(newDeviceCondition.SetPoint),
		Duration:    PointerToNullInt32(newDeviceCondition.Duration),
		Operator:    PointerToNullInt32(newDeviceCondition.Operator),
	}

	id, err := dcr.queries.CreateDeviceCondition(ctx, arg)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (dcr DeviceConditionRepository) GetDeviceConditionFromID(ID int) (*domain.DeviceCondition, error) {
	ctx := context.Background()

	deviceCondition, err := dcr.queries.GetDeviceConditionFromID(ctx, int32(ID))
	if err != nil {
		return nil, err
	}

	getDeviceCondition := domain.NewDeviceCondition(
		int(deviceCondition.ID),
		int(deviceCondition.DeviceID),
		int(deviceCondition.OperationID),
		deviceCondition.Valid,
		NullFloat64ToPointer(deviceCondition.SetPoint),
		NullInt32ToPointer(deviceCondition.Duration),
		NullInt32ToPointer(deviceCondition.Operator),
	)

	return getDeviceCondition, nil
}

func (dcr DeviceConditionRepository) GetDeviceConditionsFromDeviceID(deviceID int) ([]*domain.DeviceCondition, error) {
	ctx := context.Background()

	deviceConditionsRow, err := dcr.queries.GetDeviceConditionsFromDeviceID(ctx, int32(deviceID))
	if err != nil {
		return nil, err
	}

	deviceConditions := make([]*domain.DeviceCondition, len(deviceConditionsRow))
	for i, v := range deviceConditionsRow {
		deviceConditions[i] = domain.NewDeviceCondition(
			int(v.ID),
			int(v.DeviceID),
			int(v.OperationID),
			v.Valid,
			NullFloat64ToPointer(v.SetPoint),
			NullInt32ToPointer(v.Duration),
			NullInt32ToPointer(v.Operator),
		)
	}

	return deviceConditions, nil
}
