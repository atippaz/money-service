package repositories

import "github.com/google/uuid"

type SpendingTypeRepository interface {
	GetSpendingTypes() (*[]SpendingTypeResultQuery, error)
}
type SpendingTypeResultQuery struct {
	SpendingTypeId uuid.UUID
	NameTh         string
	NameEn         string
}
