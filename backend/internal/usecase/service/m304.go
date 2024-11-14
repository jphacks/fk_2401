package service

import "github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"

type M304Service struct {
	m304Repository M304RepositoryInterface
}

func NewM304Service(mr M304RepositoryInterface) *M304Service {
	return &M304Service{
		m304Repository: mr,
	}
}

func (ms M304Service) CreateM304(newM304 domain.M304) (int64, error) {
	id, err := ms.m304Repository.CreateM304(newM304)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ms M304Service) GetM304FromID(ID int) (*domain.M304, error) {
	m304s, err := ms.m304Repository.GetM304FromID(ID)
	if err != nil {
		return nil, err
	}

	return m304s, nil
}
