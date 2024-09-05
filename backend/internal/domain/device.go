package domain

type Device struct {
	ID            int
	HouseID       int
	ClimateDataID int
	Duration      int
}

func NewDeviceWithID(id, houseId, climateDataId, duration int) *Device {
	return &Device{
		ID:            id,
		HouseID:       houseId,
		ClimateDataID: climateDataId,
		Duration:      duration,
	}
}

func NewDevice(houseId, climateDataId, duration int) *Device {
	return &Device{
		HouseID:       houseId,
		ClimateDataID: climateDataId,
		Duration:      duration,
	}
}
