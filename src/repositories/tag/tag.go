package repositories

import (
	"github.com/google/uuid"
)

type TagRepository interface {
	GetTags(userId *uuid.UUID, hasSystem bool) (*[]TagResultQuery, error)
	CreateTag(userOwner uuid.UUID, payload TagInsertDB) (*uuid.UUID, error)
}
type TagResultQuery struct {
	TagId          uuid.UUID
	NameTh         string
	NameEn         string
	IsActive       bool
	SpendingTypeId uuid.UUID
	Owner          uuid.UUID
}
type TagInsertDB struct {
	NameTh         string
	NameEn         string
	SpendingTypeId uuid.UUID
}
