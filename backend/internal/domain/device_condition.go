package domain

type DeviceCondition struct {
	ID       int
	DeviceID int
	Valid    bool
	SetPoint *float64
	Duration *int
	Operator *int
}

func NewDeviceCondition(id int, deviceID int, valid bool, setPoint *float64, duration *int, operator *int) *DeviceCondition {
	return &DeviceCondition{
		ID:       id,
		DeviceID: deviceID,
		Valid:    valid,
		SetPoint: setPoint,
		Duration: duration,
		Operator: operator,
	}
}

func (deviceCondition *DeviceCondition) ChangeValid(valid bool) {
	deviceCondition.Valid = valid
}

func (deviceCondition *DeviceCondition) ChangeSetPoint(setPoint *float64) {
	deviceCondition.SetPoint = setPoint
}

func (deviceCondition *DeviceCondition) ChangeDuration(duration *int) {
	deviceCondition.Duration = duration
}

func (deviceCondition *DeviceCondition) ChangeOperator(operator *int) {
	deviceCondition.Operator = operator
}
