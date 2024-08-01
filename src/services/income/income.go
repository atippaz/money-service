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
		Date:  payload.Date,
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
func (s *incomeService) GetSummary(id uuid.UUID, startDate *time.Time, endDate *time.Time) (*[]IncomeSummaryResult, error) {
	res, err := s.repo.GetIncomesByUser(id, startDate, endDate)

	resultMap := make(map[string]*IncomeSummaryResult)
	for _, result := range *res {
		if existing, ok := resultMap[result.TagId.String()]; ok {
			existing.Value.Add(result.Value)
		} else {
			resultMap[result.TagId.String()] = &IncomeSummaryResult{
				TagId:          result.TagId,
				Value:          result.Value,
				SpendingTypeId: "",
			}
		}
	}

	var groupedResults []IncomeSummaryResult
	for _, result := range resultMap {
		groupedResults = append(groupedResults, *result)
	}
	return &groupedResults, err
}
