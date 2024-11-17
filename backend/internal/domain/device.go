package domain

type Device struct {
	ID            int
	HouseID       int
	ClimateDataID int
	M304ID        int
	SensorID      int
	DeviceName    *string
	Rly           *int
}

func NewDeviceWithID(id int, houseID int, climateDataID int, m304ID int, sensorID int, deviceName *string, rly *int) *Device {
	return &Device{
		ID:            id,
		HouseID:       houseID,
		ClimateDataID: climateDataID,
		M304ID:        m304ID,
		SensorID:      sensorID,
		DeviceName:    deviceName,
		Rly:           rly,
	}
}

func NewDevice(houseID int, climateDataID int, m304ID int, sensorID int, deviceName *string, rly *int) *Device {
	return &Device{
		HouseID:       houseID,
		ClimateDataID: climateDataID,
		M304ID:        m304ID,
		SensorID:      sensorID,
		DeviceName:    deviceName,
		Rly:           rly,
	}
}
