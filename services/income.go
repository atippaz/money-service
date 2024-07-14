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

func (s *IIncomeService) GetUserById(id string) (*interfaces.UserResult, error) {
	res, err := s.repo.GetUserById(id)
	return res, err
}
