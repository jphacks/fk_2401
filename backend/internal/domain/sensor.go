package domain

type Sensor struct {
	ID       int
	Ccmtype  string
	Room     int
	Region   int
	Order    int
	Priority int
}

func NewSensor(ccmtype string, room int, region int, order int, priority int) *Sensor {
	return &Sensor{
		Ccmtype:  ccmtype,
		Room:     room,
		Region:   region,
		Order:    order,
		Priority: priority,
	}
}

func NewSensorWithID(id int, ccmtype string, room int, region int, order int, priority int) *Sensor {
	return &Sensor{
		ID:       id,
		Ccmtype:  ccmtype,
		Room:     room,
		Region:   region,
		Order:    order,
		Priority: priority,
	}
}
