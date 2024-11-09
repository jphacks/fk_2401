package repository

import (
	"context"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
)

type OperationRepository struct {
	queries *mysqlc.Queries
}

func NewOperation(queries *mysqlc.Queries) *OperationRepository {
	return &OperationRepository{
		queries: queries,
	}
}

func (or OperationRepository) CreateOperation(newOperation domain.Operation) (int64, error) {
	ctx := context.Background()

	arg := mysqlc.CreateOperationParams{
		DeviceID: int32(newOperation.DeviceID),
		Name:     newOperation.Name,
		RlyOn:    int32(newOperation.RlyOn),
	}

	id, err := or.queries.CreateOperation(ctx, arg)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (or OperationRepository) GetOperationFromID(ID int) (*domain.Operation, error) {
	ctx := context.Background()

	operation, err := or.queries.GetOperationFromID(ctx, int32(ID))
	if err != nil {
		return nil, err
	}

	getOperation := domain.NewOperation(
		int(operation.ID),
		int(operation.DeviceID),
		operation.Name,
		int(operation.RlyOn),
	)

	return getOperation, nil
}

func (or OperationRepository) GetOperationsFromDeviceID(deviceID int) ([]*domain.Operation, error) {
	ctx := context.Background()

	operationsRow, err := or.queries.GetOperationsFromDeviceID(ctx, int32(deviceID))
	if err != nil {
		return nil, err
	}

	operations := make([]*domain.Operation, len(operationsRow))
	for i, v := range operationsRow {
		operations[i] = domain.NewOperation(
			int(v.ID),
			int(v.DeviceID),
			v.Name,
			int(v.RlyOn),
		)
	}

	return operations, nil
}
