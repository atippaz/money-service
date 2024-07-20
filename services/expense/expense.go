package services

import (
	repositories "money-service/repositories/expense"

	"github.com/google/uuid"
)

func NewExpenseService(repo repositories.ExpenseRepository) ExpenseService {
	return &expenseService{repo: repo}
}

func (s *expenseService) GetExpensesByUser(userId uuid.UUID) (*[]ExpenseResult, error) {
	// res, err := s.repo.GetExpensesByUser(userId)
	// return res, err
	return nil, nil
}
func (s *expenseService) CreateExpense(userId uuid.UUID, payload ExpenseInsert) (*uuid.UUID, error) {
	// res, err := s.repo.CreateExpense(userId, payload)
	// return res, err
	return nil, nil
}
