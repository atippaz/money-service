package services

import (
	"money-service/interfaces"
	repositories "money-service/repositories/income"

	"github.com/google/uuid"
)

type IIncomeService struct {
	repo repositories.IIncomeRepository
}

func IncomeService(repo repositories.IIncomeRepository) *IIncomeService {
	return &IIncomeService{repo: repo}
}
func (s *IIncomeService) CreateIncome(id uuid.UUID, payload interfaces.IncomeInsertDb) (*uuid.UUID, error) {
	res, err := s.repo.CreateIncome(id, payload)
	return res, err
}
func (s *IIncomeService) GetIncomesByUser(id uuid.UUID) (*[]interfaces.IncomeResultQuery, error) {
	res, err := s.repo.GetIncomesByUser(id)
	return res, err
}
