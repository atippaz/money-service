package repositories

import "github.com/google/uuid"

type SystemTagRepository interface {
	GetAllSystemTags() (*[]SystemTagResultQuery, error)
}
type SystemTagResultQuery struct {
	TagId          uuid.UUID
	NameTh         string
	NameEn         string
	IsActive       bool
	SpendingTypeId uuid.UUID
}
