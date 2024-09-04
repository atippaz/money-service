package services

import (
	user_share_repo "money-service/src/repositories/user_share"

	"github.com/google/uuid"
)

type UserShareService interface {
	Insert(userOwner uuid.UUID, payload user_share_repo.UserShareInsertDB) (*[]user_share_repo.UserShareResultQuery, error)
	GetAll(userId *uuid.UUID) (*[]user_share_repo.UserShareResultQuery, error)
}
type userShareService struct {
	repo user_share_repo.UserShareRepository
}
type UserShareResult struct {
	SpendingTypeId uuid.UUID
	NameTh         string
	NameEn         string
}
