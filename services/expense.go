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

func (s *IExpenseService) GetExpensesByUser(userId string) (*[]interfaces.ExpenseResultQuery, error) {
	res, err := s.repo.GetExpensesByUser(userId)
	return res, err
}
