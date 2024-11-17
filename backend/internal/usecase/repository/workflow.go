package repository

import "github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"

type WorkflowRepository struct {
	queries *mysqlc.Queries
}

func NewWorkflowRepository(queries *mysqlc.Queries) *ClimateDataRepository {
	return &ClimateDataRepository{
		queries: queries,
	}
}

func (wr WorkflowRepository) GetAllWorkflows()
