package services

import (
	repositories "money-service/src/repositories/expense"
	"time"

	"github.com/google/uuid"
)

func NewExpenseService(repo repositories.ExpenseRepository) ExpenseService {
	return &expenseService{repo: repo}
}

func (s *expenseService) GetExpensesByUser(userId uuid.UUID, startDate *time.Time, endDate *time.Time) (*[]ExpenseResult, error) {
	res, err := s.repo.GetExpensesByUser(userId, startDate, endDate)
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
func (s *expenseService) GetSummary(id uuid.UUID, startDate *time.Time, endDate *time.Time) (*[]ExpenseSummaryResult, error) {
	res, err := s.repo.GetExpensesByUser(id, startDate, endDate)
	resultMap := make(map[string]*ExpenseSummaryResult)
	for _, result := range *res {
		if existing, ok := resultMap[result.TagId.String()]; ok {
			existing.Value.Add(result.Value)
		} else {
			resultMap[result.TagId.String()] = &ExpenseSummaryResult{
				TagId: result.TagId,
				Value: result.Value,
			}
		}
	}

	var groupedResults []ExpenseSummaryResult
	for _, result := range resultMap {
		groupedResults = append(groupedResults, *result)
	}
	return &groupedResults, err
}
