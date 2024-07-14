package services

import (
	"money-service/interfaces"
	repositories "money-service/repositories/income"
)

type IIncomeService struct {
	repo repositories.IIncomeRepository
}

func IncomeService(repo repositories.IIncomeRepository) *IIncomeService {
	return &IIncomeService{repo: repo}
}

func (s *IIncomeService) GetIncomesByUser(id string) (*[]interfaces.IncomeResultQuery, error) {
	res, err := s.repo.GetIncomesByUser(id)
	return res, err
}
