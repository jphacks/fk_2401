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

func (cds ClimateDataService) GetClimateData() ([]*domain.ClimateData, error) {
	climateData, err := cds.climateDataRepository.GetAllClimateData()
	if err != nil {
		return nil, err
	}

	return climateData, nil
}
