package services

import (
	"money-service/interfaces"
	repositories "money-service/repositories/income"

	"github.com/google/uuid"
)

type IncomeService struct {
	repo repositories.IncomeRepository
}

func NewIncomeService(repo repositories.IncomeRepository) *IncomeService {
	return &IncomeService{repo: repo}
}
func (s *IncomeService) CreateIncome(id uuid.UUID, payload interfaces.IncomeInsertDb) (*uuid.UUID, error) {
	res, err := s.repo.CreateIncome(id, payload)
	return res, err
}
func (s *IncomeService) GetIncomesByUser(id uuid.UUID) (*[]interfaces.IncomeResultQuery, error) {
	res, err := s.repo.GetIncomesByUser(id)
	return res, err
}
