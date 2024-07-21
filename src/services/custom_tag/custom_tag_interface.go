package services

import (
	repositories "money-service/src/repositories/custom_tag"

	"github.com/google/uuid"
)

type CustomTagService interface {
	GetCustomTagsByUser(userId uuid.UUID) (*[]CustomTagResult, error)
	CreateCustomTag(id uuid.UUID) (*uuid.UUID, error)
}
type customTagService struct {
	repo repositories.CustomTagRepository
}
type CustomTagResult struct {
	TagId          uuid.UUID
	NameTh         string
	NameEn         string
	IsActive       bool
	SpendingTypeId uuid.UUID
	UserOwner      uuid.UUID
}
type CustomTagInsert struct {
	NameTh         string
	NameEn         string
	SpendingTypeId uuid.UUID
}
