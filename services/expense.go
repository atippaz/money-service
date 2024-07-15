package services

import (
	"money-service/interfaces"
	repositories "money-service/repositories/expense"

	"github.com/google/uuid"
)

type IExpenseService struct {
	repo repositories.IExpenseRepository
}

func ExpenseService(repo repositories.IExpenseRepository) *IExpenseService {
	return &IExpenseService{repo: repo}
}

func (s *IExpenseService) GetExpensesByUser(userId uuid.UUID) (*[]interfaces.ExpenseResultQuery, error) {
	res, err := s.repo.GetExpensesByUser(userId)
	return res, err
}
func (s *IExpenseService) CreateExpense(userId uuid.UUID, payload interfaces.ExpenseInsertDb) (*uuid.UUID, error) {
	res, err := s.repo.CreateExpense(userId, payload)
	return res, err
}
