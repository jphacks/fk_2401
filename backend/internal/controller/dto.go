package controller

type CreateHouseController struct {
	Name string `json:"name" binding:"required,min=1,max=12,alphanum"`
}

type CreateDeviceController struct {
	ClimateDataID int     `json:"climate_data_id" binding:"required,number"`
	DeviceName    string  `json:"device_name" binding:"required,min=1,max=12"`
	SetPoint      float64 `json:"set_point" binding:"number"`
	Duration      int     `json:"duration" binding:"number"`
}

type HouseResponse struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type DeviceResponse struct {
	ID          int     `json:"id,omitempty"`
	HouseID     int     `json:"house_id,omitempty"`
	Name        string  `json:"name,omitempty"`
	SetPoint    float64 `json:"set_point,omitempty"`
	Duration    int     `json:"duration,omitempty"`
	ClimateData string  `json:"climate_data,omitempty"`
	Unit        string  `json:"unit,omitempty"`
}

type ClimateDataResponse struct {
	ID          int    `json:"id,omitempty"`
	ClimateData string `json:"climate_data,omitempty"`
	Unit        string `json:"unit,omitempty"`
}

func NewHouseResponse(id int, name string) *HouseResponse {
	return &HouseResponse{
		ID:   id,
		Name: name,
	}
}

func NewDeviceResponse(
	id int,
	houseID int,
	deviceName string,
	setPoint float64,
	duration int,
	climateData string,
	unit string,
) *DeviceResponse {
	return &DeviceResponse{
		ID:          id,
		HouseID:     houseID,
		Name:        deviceName,
		SetPoint:    setPoint,
		Duration:    duration,
		ClimateData: climateData,
		Unit:        unit,
	}
}

func NewClimateDataResponse(id int, climateData, unit string) *ClimateDataResponse {
	return &ClimateDataResponse{
		ID:          id,
		ClimateData: climateData,
		Unit:        unit,
	}
}
