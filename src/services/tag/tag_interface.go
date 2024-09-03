package services

import (
	repositories "money-service/src/repositories/tag"

	"github.com/google/uuid"
)

type TagService interface {
	GetTagsByUser(userId uuid.UUID, hasSystem bool) (*[]TagResult, error)
	GetSystemTags() (*[]TagResult, error)
	CreateTag(id uuid.UUID) (*uuid.UUID, error)
}
type tagService struct {
	repo repositories.TagRepository
}
type TagResult struct {
	TagId          uuid.UUID
	NameTh         string
	NameEn         string
	IsActive       bool
	SpendingTypeId uuid.UUID
	Owner          uuid.UUID
}
type TagInsert struct {
	NameTh         string
	NameEn         string
	SpendingTypeId uuid.UUID
}
