package service

import "github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"

type M304RecordService struct {
	m304RecordRepository M304RecordRepositoryInterface
}

func NewM304RecordService(mrr M304RecordRepositoryInterface) *M304RecordService {
	return &M304RecordService{
		m304RecordRepository: mrr,
	}
}

func (mrs M304RecordService) CreateM304Record(newM304Record domain.M304Record) (int64, error) {
	id, err := mrs.m304RecordRepository.CreateM304Record(newM304Record)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (mrs M304RecordService) GetM304RecordFromM304ID(m304ID int) ([]*domain.M304Record, error) {
	m304Records, err := mrs.m304RecordRepository.GetM304RecordFromM304ID(m304ID)
	if err != nil {
		return nil, err
	}

	return m304Records, nil
}
