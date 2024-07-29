package services

import (
	repositories "money-service/src/repositories/income"
	"time"

	"github.com/google/uuid"
)

func NewIncomeService(repo repositories.IncomeRepository) IncomeService {
	return &incomeService{repo: repo}
}
func (s *incomeService) CreateIncome(id uuid.UUID, payload IncomeInsert) (*uuid.UUID, error) {
	res, err := s.repo.CreateIncome(id, repositories.IncomeInsertDb{
		TagId: payload.TagId,
		Value: payload.Value,
	})
	return res, err
}
func (s *incomeService) GetIncomesByUser(id uuid.UUID, startDate *time.Time, endDate *time.Time) (*[]IncomeResult, error) {
	res, err := s.repo.GetIncomesByUser(id, startDate, endDate)
	var results []IncomeResult
	for _, result := range *res {
		results = append(results, IncomeResult{
			IncomeId:    result.IncomeId,
			CreatedDate: result.CreatedDate,
			TagId:       result.TagId,
			Value:       result.Value,
			UserOwner:   result.UserOwner,
		})
	}
	return &results, err
}
