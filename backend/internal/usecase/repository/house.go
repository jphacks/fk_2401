package repository

import (
	"context"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
)

type HouseRepository struct {
	queries *mysqlc.Queries
}

func NewHouseRepository(queries *mysqlc.Queries) *HouseRepository {
	return &HouseRepository{
		queries: queries,
	}
}

func (hr HouseRepository) GetAllHouse() ([]*domain.House, error) {
	ctx := context.Background()

	housesRow, err := hr.queries.GetAllHouse(ctx)
	if err != nil {
		return nil, err
	}
	houses := make([]*domain.House, len(housesRow))
	for i, v := range housesRow {
		houses[i] = domain.NewHouseWithID(
			int(v.ID),
			v.Name,
		)
	}

	return houses, nil
}
