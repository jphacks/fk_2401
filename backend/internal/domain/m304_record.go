package domain

type M304Record struct {
	ID                int
	M304ID            int
	DeviceConditionID int
	Block             string
	Valid             bool
	Position          int
}

func NewM304Record(id int, m304ID int, deviceConditionID int, block string, valid bool, position int) *M304Record {
	return &M304Record{
		ID:                id,
		M304ID:            m304ID,
		DeviceConditionID: deviceConditionID,
		Block:             block,
		Valid:             valid,
		Position:          position,
	}
}
