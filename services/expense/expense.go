package services

import (
	repositories "money-service/repositories/expense"

	"github.com/google/uuid"
)

func NewExpenseService(repo repositories.ExpenseRepository) ExpenseService {
	return &expenseService{repo: repo}
}

func (s *expenseService) GetExpensesByUser(userId uuid.UUID) (*[]ExpenseResult, error) {
	res, err := s.repo.GetExpensesByUser(userId)
	var results []ExpenseResult
	for _, result := range *res {
		results = append(results, ExpenseResult{
			ExpenseId:   result.ExpenseId,
			CreatedDate: result.CreatedDate,
			TagId:       result.TagId,
			Value:       result.Value,
			UserOwner:   result.UserOwner,
		})
	}
	return &results, err
}
func (s *expenseService) CreateExpense(userId uuid.UUID, payload ExpenseInsert) (*uuid.UUID, error) {
	res, err := s.repo.CreateExpense(userId, repositories.ExpenseInsertDb{
		TagId: payload.TagId,
		Value: payload.Value,
	})
	return res, err
}
