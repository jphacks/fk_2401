package domain

type Device struct {
	ID            int
	HouseID       int
	ClimateDataID int
	SetPoint      float64
	Duration      int
}

func NewDeviceWithID(id int, houseID int, climateDataID int, setPoint float64, duration int) *Device {
	return &Device{
		ID:            id,
		HouseID:       houseID,
		ClimateDataID: climateDataID,
		SetPoint:      setPoint,
		Duration:      duration,
	}
}

func NewDevice(houseID int, climateDataID int, setPoint float64, duration int) *Device {
	return &Device{
		HouseID:       houseID,
		ClimateDataID: climateDataID,
		SetPoint:      setPoint,
		Duration:      duration,
	}
}

func (device *Device) ChangeSetPoint(setPoint float64) {
	device.SetPoint = setPoint
}

func (device *Device) ChangeDuration(duration int) {
	device.Duration = duration
}
