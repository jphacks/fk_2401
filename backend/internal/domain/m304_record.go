package domain

type M304Record struct {
	ID       int
	M304ID   int
	DeviceID int
	Block    string
	Valid    bool
	Position int
}

func NewM304Record(id int, m304ID int, deviceID int, block string, valid bool, position int) *M304Record {
	return &M304Record{
		ID:       id,
		M304ID:   m304ID,
		DeviceID: deviceID,
		Block:    block,
		Valid:    valid,
		Position: position,
	}
}
