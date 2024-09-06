package domain

type Device struct {
	ID            int
	HouseID       int
	ClimateDataID int
	Duration      int
}

func NewDeviceWithID(id, houseID, climateDataID, duration int) *Device {
	return &Device{
		ID:            id,
		HouseID:       houseID,
		ClimateDataID: climateDataID,
		Duration:      duration,
	}
}

func NewDevice(houseID, climateDataID, duration int) *Device {
	return &Device{
		HouseID:       houseID,
		ClimateDataID: climateDataID,
		Duration:      duration,
	}
}
