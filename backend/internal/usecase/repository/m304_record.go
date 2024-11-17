package repository

import (
	"context"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
)

type M304RecordRepository struct {
	queries *mysqlc.Queries
}

func NewM304RecordRepository(queries *mysqlc.Queries) *M304RecordRepository {
	return &M304RecordRepository{
		queries: queries,
	}
}

func (mrr M304RecordRepository) CreateM304Record(NewM304Record domain.M304Record) (int64, error) {
	ctx := context.Background()

	arg := mysqlc.CreateM304RecordParams{
		M304ID:            int32(NewM304Record.M304ID),
		DeviceConditionID: int32(NewM304Record.DeviceConditionID),
		Block:             NewM304Record.Block,
		Valid:             NewM304Record.Valid,
		Position:          int32(NewM304Record.Position),
	}

	id, err := mrr.queries.CreateM304Record(ctx, arg)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (mrr M304RecordRepository) GetM304RecordFromM304ID(m304ID int) ([]*domain.M304Record, error) {
	ctx := context.Background()

	m304RecordRow, err := mrr.queries.GetRecordFromM304ID(ctx, int32(m304ID))
	if err != nil {
		return nil, err
	}

	m304Records := make([]*domain.M304Record, len(m304RecordRow))
	for i, v := range m304RecordRow {
		m304Records[i] = domain.NewM304Record(
			int(v.ID),
			int(v.M304ID),
			int(v.DeviceConditionID),
			v.Block,
			v.Valid,
			int(v.Position),
		)
	}

	return m304Records, nil
}
