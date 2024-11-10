package service

import "github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"

type ClimateDataService struct {
	climateDataRepository ClimateDataRepositoryInterface
}

func NewClimateDataService(cdr ClimateDataRepositoryInterface) *ClimateDataService {
	return &ClimateDataService{
		climateDataRepository: cdr,
	}
}

func (cds ClimateDataService) GetAllClimateData() ([]*domain.ClimateData, error) {
	climateData, err := cds.climateDataRepository.GetAllClimateData()
	if err != nil {
		return nil, err
	}

	return climateData, nil
}

func (cds ClimateDataService) GetClimateDataFromID(ID int) (*domain.ClimateData, error) {
	climateData, err := cds.climateDataRepository.GetClimateDataFromID(ID)
	if err != nil {
		return nil, err
	}

	return climateData, nil
}
