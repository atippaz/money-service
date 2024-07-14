package services

import (
	"money-service/interfaces"
	repositories "money-service/repositories/spending_type"
)

type ISpendingTypeService struct {
	repo repositories.ISpendingTypeRepository
}

func SpendingTypeService(repo repositories.ISpendingTypeRepository) *ISpendingTypeService {
	return &ISpendingTypeService{repo: repo}
}

func (s *ISpendingTypeService) GetSpendingTypes() (*[]interfaces.SpendingTypeResult, error) {
	res, err := s.repo.GetSpendingTypes()
	return res, err
}
