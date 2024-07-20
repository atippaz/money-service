package services

import (
	spending_type_repo "money-service/repositories/spending_type"

	"github.com/google/uuid"
)

type SpendingTypeService interface {
	GetSpendingTypes() (*[]SpendingTypeResult, error)
}
type spendingTypeService struct {
	repo spending_type_repo.SpendingTypeRepository
}
type SpendingTypeResult struct {
	SpendingTypeId uuid.UUID
	NameTh         string
	NameEn         string
}
