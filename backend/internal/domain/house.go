package domain

type House struct {
	ID   int
	Name string
}

func NewHouseWithID(id int, name string) *House {
	return &House{
		ID:   id,
		Name: name,
	}
}

func NewHouse(name string) *House {
	return &House{
		Name: name,
	}
}
