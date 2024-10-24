package domain

type Device struct {
	ID            int
	HouseID       int
	ClimateDataID int
	UecsDeviceID  int
	DeviceName    *string
	Valid         *bool
	SetPoint      *float64
	Duration      *int
	Operator      *int
}

func NewDeviceWithID(id int, houseID int, climateDataID int, uecsDeviceID int, deviceName *string, valid *bool, setPoint *float64, duration *int, operator *int) *Device {
	return &Device{
		ID:            id,
		HouseID:       houseID,
		ClimateDataID: climateDataID,
		UecsDeviceID:  uecsDeviceID,
		DeviceName:    deviceName,
		Valid:         valid,
		SetPoint:      setPoint,
		Duration:      duration,
		Operator:      operator,
	}
}

func NewDevice(houseID int, climateDataID int, uecsDeviceID int, deviceName *string, valid *bool, setPoint *float64, duration *int, operator *int) *Device {
	return &Device{
		HouseID:       houseID,
		ClimateDataID: climateDataID,
		UecsDeviceID:  uecsDeviceID,
		DeviceName:    deviceName,
		Valid:         valid,
		SetPoint:      setPoint,
		Duration:      duration,
		Operator:      operator,
	}
}

func (device *Device) ChangeValid(valid *bool) {
	device.Valid = valid
}

func (device *Device) ChangeSetPoint(setPoint *float64) {
	device.SetPoint = setPoint
}

func (device *Device) ChangeDuration(duration *int) {
	device.Duration = duration
}

func (device *Device) ChangeOperator(operator *int) {
	device.Operator = operator
}
