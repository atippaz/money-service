package services

import (
	spending_type_repo "money-service/src/repositories/spending_type"
)

func NewSpendingTypeService(repo spending_type_repo.SpendingTypeRepository) SpendingTypeService {
	return &spendingTypeService{repo: repo}
}

func (s *spendingTypeService) GetSpendingTypes() (*[]SpendingTypeResult, error) {
	res, err := s.repo.GetSpendingTypes()
	var results []SpendingTypeResult
	for _, data := range *res {
		results = append(results, SpendingTypeResult{
			SpendingTypeId: data.SpendingTypeId,
			NameTh:         data.NameTh,
			NameEn:         data.NameEn,
		})
	}
	return &results, err
}
