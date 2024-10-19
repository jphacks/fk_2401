package domain

type UecsDevice struct {
	ID       int
	Ccmtype  string
	Room     int
	Region   int
	Order    int
	Priority int
}

func NewUecsDevice(ccmtype string, room int, region int, order int, priority int) *UecsDevice {
	return &UecsDevice{
		Ccmtype:  ccmtype,
		Room:     room,
		Region:   region,
		Order:    order,
		Priority: priority,
	}
}

func NewUecsDeviceWithID(id int, ccmtype string, room int, region int, order int, priority int) *UecsDevice {
	return &UecsDevice{
		ID:       id,
		Ccmtype:  ccmtype,
		Room:     room,
		Region:   region,
		Order:    order,
		Priority: priority,
	}
}
