package service

import "github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"

type OperationService struct {
	operationRepository OperationRepositoryInterface
}

func NewOperationService(or OperationRepositoryInterface) *OperationService {
	return &OperationService{
		operationRepository: or,
	}
}

func (os OperationService) CreateOperation(newOperation domain.Operation) (int64, error) {
	id, err := os.operationRepository.CreateOperation(newOperation)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (os OperationService) GetOperationFromID(ID int) (*domain.Operation, error) {
	operation, err := os.operationRepository.GetOperationFromID(ID)
	if err != nil {
		return nil, err
	}

	return operation, nil
}

func (os OperationService) GetOperationsFromDeviceID(deviceID int) ([]*domain.Operation, error) {
	operations, err := os.operationRepository.GetOperationsFromDeviceID(deviceID)
	if err != nil {
		return nil, err
	}

	return operations, nil
}
