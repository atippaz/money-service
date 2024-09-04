package services

import (
	share_repo "money-service/src/repositories/share"

	"github.com/google/uuid"
)

type ShareService interface {
	Insert(userOwner uuid.UUID, payload share_repo.ShareInsertDB) (*[]share_repo.ShareResultQuery, error)
	GetAll(userId *uuid.UUID) (*[]share_repo.ShareResultQuery, error)
	GetById(shareId *uuid.UUID) (*[]share_repo.ShareResultQuery, error)
}
type shareService struct {
	repo share_repo.ShareRepository
}
type ShareResult struct {
	SpendingTypeId uuid.UUID
	NameTh         string
	NameEn         string
}
