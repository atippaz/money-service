package repositories

import "github.com/google/uuid"

type UserShareRepository interface {
	Insert(userOwner uuid.UUID, payload UserShareInsertDB) (*[]UserShareResultQuery, error)
	GetAll(userId *uuid.UUID) (*[]UserShareResultQuery, error)
}
type UserShareResultQuery struct {
	SpendingTypeId uuid.UUID
	NameTh         string
	NameEn         string
}

type UserShareInsertDB struct {
	NameTh         string
	NameEn         string
	SpendingTypeId uuid.UUID
}
