package mocks

import "github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"

type MockClimateDataRepository struct {
	ClimateDataList []*domain.ClimateData
}

func NewMockClimateDataRepository() *MockClimateDataRepository {
	climateDataList := make([]*domain.ClimateData, 4)
	climateDataList[0] = &domain.ClimateData{
		ID:          1,
		ClimateData: "気温",
		Unit:        "℃",
	}
	climateDataList[1] = &domain.ClimateData{
		ID:          2,
		ClimateData: "湿度",
		Unit:        "%",
	}
	climateDataList[2] = &domain.ClimateData{
		ID:          3,
		ClimateData: "二酸化炭素量",
		Unit:        "ppm",
	}
	climateDataList[3] = &domain.ClimateData{
		ID:          4,
		ClimateData: "照度",
		Unit:        "lx",
	}

	return &MockClimateDataRepository{
		ClimateDataList: climateDataList,
	}
}

func (cdr MockClimateDataRepository) GetAllClimateData() ([]*domain.ClimateData, error) {
	return cdr.ClimateDataList, nil
}
