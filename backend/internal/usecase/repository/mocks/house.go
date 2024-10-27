package mocks

import "github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"

type MockHouseRepository struct {
	HousesList []*domain.House
}

func NewMockHouseRepository() *MockHouseRepository {
	housesList := make([]*domain.House, 4)

	housesList[0] = &domain.House{
		ID:   1,
		Name: "aaaaa",
	}
	housesList[1] = &domain.House{
		ID:   2,
		Name: "bbbbb",
	}
	housesList[2] = &domain.House{
		ID:   3,
		Name: "ccccc",
	}
	housesList[3] = &domain.House{
		ID:   4,
		Name: "ddddd",
	}

	return &MockHouseRepository{
		HousesList: housesList,
	}
}

func (hr MockHouseRepository) GetAllHouses() ([]*domain.House, error) {
	return hr.HousesList, nil
}

func (hr MockHouseRepository) CreateHouse(name string) (int64, error) {
	return 0, nil
}
