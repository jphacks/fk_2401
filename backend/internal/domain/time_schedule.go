package domain

type TimeSchedule struct {
	ID                int
	DeviceConditionID int
	StartTime         string
	EndTime           string
}

func NewTimeSchedule(id int, deviceConditionID int, startTime string, endTime string) *TimeSchedule {
	return &TimeSchedule{
		ID:                id,
		DeviceConditionID: deviceConditionID,
		StartTime:         startTime,
		EndTime:           endTime,
	}
}
