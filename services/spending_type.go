package services

import (
	"money-service/interfaces"
	repositories "money-service/repositories/spending_type"
)

type SpendingTypeService struct {
	repo repositories.SpendingTypeRepository
}

func NewSpendingTypeService(repo repositories.SpendingTypeRepository) *SpendingTypeService {
	return &SpendingTypeService{repo: repo}
}

func (s *SpendingTypeService) GetSpendingTypes() (*[]interfaces.SpendingTypeResultResponse, error) {
	res, err := s.repo.GetSpendingTypes()
	var results []interfaces.SpendingTypeResultResponse
	for _, data := range *res {
		results = append(results, interfaces.SpendingTypeResultResponse{
			SpendingTypeId: data.SpendingTypeId,
			NameTh:         data.NameTh,
			NameEn:         data.NameEn,
		})
	}
	return &results, err
}
