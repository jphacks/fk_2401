package repository

import (
	"context"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
)

type TimeScheduleRepository struct {
	queries *mysqlc.Queries
}

func NewTimeScheduleRepository(queries *mysqlc.Queries) *TimeScheduleRepository {
	return &TimeScheduleRepository{
		queries: queries,
	}
}

func (tsr TimeScheduleRepository) CreateTimeSchedule(newTimeSchedule domain.TimeSchedule) (int64, error) {
	ctx := context.Background()

	arg := mysqlc.CreateTimeScheduleParams{
		DeviceConditionID: int32(newTimeSchedule.DeviceConditionID),
		StartTime:         newTimeSchedule.StartTime,
		EndTime:           newTimeSchedule.EndTime,
	}

	id, err := tsr.queries.CreateTimeSchedule(ctx, arg)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (tsr TimeScheduleRepository) GetTimeScheduleFromID(ID int) (*domain.TimeSchedule, error) {
	ctx := context.Background()

	timeSchedule, err := tsr.queries.GetTimeScheduleFromID(ctx, int32(ID))
	if err != nil {
		return nil, err
	}

	getTimeSchedule := domain.NewTimeSchedule(
		int(timeSchedule.ID),
		int(timeSchedule.DeviceConditionID),
		timeSchedule.StartTime,
		timeSchedule.EndTime,
	)

	return getTimeSchedule, nil
}

func (tsr TimeScheduleRepository) GetTimeSchedulesFromDeviceCondition(deviceConditionID int) ([]*domain.TimeSchedule, error) {
	ctx := context.Background()

	timeSchedulesRow, err := tsr.queries.GetTimeSchedulesFromDeviceCondition(ctx, int32(deviceConditionID))
	if err != nil {
		return nil, err
	}

	timeSchedules := make([]*domain.TimeSchedule, len(timeSchedulesRow))
	for i, v := range timeSchedulesRow {
		timeSchedules[i] = domain.NewTimeSchedule(
			int(v.ID),
			int(v.DeviceConditionID),
			v.StartTime,
			v.EndTime,
		)
	}

	return timeSchedules, nil
}
