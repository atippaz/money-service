package services

import (
	"money-service/interfaces"
	repositories "money-service/repositories/expense"

	"github.com/google/uuid"
)

type ExpenseService struct {
	repo repositories.ExpenseRepository
}

func NewExpenseService(repo repositories.ExpenseRepository) *ExpenseService {
	return &ExpenseService{repo: repo}
}

func (s *ExpenseService) GetExpensesByUser(userId uuid.UUID) (*[]interfaces.ExpenseResultQuery, error) {
	res, err := s.repo.GetExpensesByUser(userId)
	return res, err
}
func (s *ExpenseService) CreateExpense(userId uuid.UUID, payload interfaces.ExpenseInsertDb) (*uuid.UUID, error) {
	res, err := s.repo.CreateExpense(userId, payload)
	return res, err
}
