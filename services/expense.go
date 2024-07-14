package services

import (
	"money-service/interfaces"
	repositories "money-service/repositories/expense"
)

type IExpenseService struct {
	repo repositories.IExpenseRepository
}

func ExpenseService(repo repositories.IExpenseRepository) *IExpenseService {
	return &IExpenseService{repo: repo}
}

func (s *IExpenseService) GetUserById(id string) (*interfaces.UserResult, error) {
	res, err := s.repo.GetUserById(id)
	return res, err
}
