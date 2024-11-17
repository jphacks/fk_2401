package domain

type Operation struct {
	ID       int
	DeviceID int
	Name     string
	RlyOn    int
}

func NewOperation(id int, deviceID int, name string, rlyOn int) *Operation {
	return &Operation{
		ID:       id,
		DeviceID: deviceID,
		Name:     name,
		RlyOn:    rlyOn,
	}
}
