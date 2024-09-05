package domain

type ClimateData struct {
	ID          int
	ClimateData string
	Unit        string
}

func NewClimateData(id int, climateData, unit string) *ClimateData {
	return &ClimateData{
		ID:          id,
		ClimateData: climateData,
		Unit:        unit,
	}
}
