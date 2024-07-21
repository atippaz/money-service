package repositories

import (
	"github.com/google/uuid"
)

type CustomTagRepository interface {
	GetCustomTagsByUser(userId uuid.UUID) (*[]CustomTagResultQuery, error)
	CreateCustomTag(userOwner uuid.UUID, payload CustomTagInsertDB) (*uuid.UUID, error)
}
type CustomTagResultQuery struct {
	TagId          uuid.UUID
	NameTh         string
	NameEn         string
	IsActive       bool
	SpendingTypeId uuid.UUID
	UserOwner      uuid.UUID
}
type CustomTagInsertDB struct {
	NameTh         string
	NameEn         string
	SpendingTypeId uuid.UUID
}
