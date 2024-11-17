package service

import "github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"

type TimeScheduleService struct {
	timeScheduleRepository TimeScheduleRepositoryInterface
}

func NewTimeScheduleService(tsr TimeScheduleRepositoryInterface) *TimeScheduleService {
	return &TimeScheduleService{
		timeScheduleRepository: tsr,
	}
}

func (tss TimeScheduleService) CreateTimeSchedule(newTimeSchedule domain.TimeSchedule) (int64, error) {
	id, err := tss.timeScheduleRepository.CreateTimeSchedule(newTimeSchedule)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (tss TimeScheduleService) GetTimeScheduleFromID(ID int) (*domain.TimeSchedule, error) {
	timeSchedule, err := tss.timeScheduleRepository.GetTimeScheduleFromID(ID)
	if err != nil {
		return nil, err
	}

	return timeSchedule, nil
}

func (tss TimeScheduleService) GetTimeSchedulesFromDeviceCondition(deviceConditionID int) ([]*domain.TimeSchedule, error) {
	timeSchedules, err := tss.timeScheduleRepository.GetTimeSchedulesFromDeviceCondition(deviceConditionID)
	if err != nil {
		return nil, err
	}

	return timeSchedules, nil
}
