package controller

type CreateHouseController struct {
	Name string `json:"name" binding:"required,min=1,max=12,alphanum"`
}

type CreateDeviceController struct {
	ClimateDataID int     `json:"climate_data_id" binding:"required,number"`
	DeviceName    string  `json:"device_name" binding:"min=1,max=12"`
	SetPoint      float64 `json:"set_point" binding:"number"`
	Duration      int     `json:"duration" binding:"number"`
}

type ClimateDatasResponse struct {
	ClimateData string `json:"climate_data,omitempty"`
	Id          int    `json:"id,omitempty"`
	Unit        string `json:"unit,omitempty"`
}
