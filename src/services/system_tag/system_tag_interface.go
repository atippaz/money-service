package services

import (
	repositories "money-service/src/repositories/system_tag"

	"github.com/google/uuid"
)

type SystemTagService interface {
	GetAllSystemTags() (*[]SystemTagResult, error)
}
type systemTagService struct {
	repo repositories.SystemTagRepository
}
type SystemTagResult struct {
	TagId          uuid.UUID
	NameTh         string
	NameEn         string
	IsActive       bool
	SpendingTypeId uuid.UUID
}
