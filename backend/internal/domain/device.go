package domain

type Device struct {
	ID            int
	HouseID       int
	ClimateDataID int
	SetPoint      int
	Duration      int
}

func NewDeviceWithID(id, houseID, climateDataID, setPoint, duration int) *Device {
	return &Device{
		ID:            id,
		HouseID:       houseID,
		ClimateDataID: climateDataID,
		SetPoint:      setPoint,
		Duration:      duration,
	}
}

func NewDevice(houseID, climateDataID, setPoint, duration int) *Device {
	return &Device{
		HouseID:       houseID,
		ClimateDataID: climateDataID,
		SetPoint:      setPoint,
		Duration:      duration,
	}
}

func (device *Device) ChangeSetPoint(setPoint int) {
	device.SetPoint = setPoint
}

func (device *Device) ChangeDuration(duration int) {
	device.Duration = duration
}
