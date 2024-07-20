package services

import (
	repositories "money-service/repositories/income"

	"github.com/google/uuid"
)

func NewIncomeService(repo repositories.IncomeRepository) IncomeService {
	return &incomeService{repo: repo}
}
func (s *incomeService) CreateIncome(id uuid.UUID, payload IncomeInsert) (*uuid.UUID, error) {
	// res, err := s.repo.CreateIncome(id, payload)
	// return res, err
	return nil, nil
}
func (s *incomeService) GetIncomesByUser(id uuid.UUID) (*[]IncomeResult, error) {
	// res, err := s.repo.GetIncomesByUser(id)
	// return res, err
	return nil, nil
}
