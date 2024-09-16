package service

import (
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
)

type HouseService struct {
	houseRepository HouseRepositoryInterface
}

func NewHouseService(hr HouseRepositoryInterface) *HouseService {
	return &HouseService{
		houseRepository: hr,
	}
}

func (hr HouseService) GetHouses() ([]*domain.House, error) {
	houses, err := hr.houseRepository.GetAllHouses()
	if err != nil {
		return nil, err
	}

	return houses, err
}
