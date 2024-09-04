package repositories

import "github.com/google/uuid"

type ShareRepository interface {
	Insert(userOwner uuid.UUID, payload ShareInsertDB) (*[]ShareResultQuery, error)
	GetAll(userId *uuid.UUID) (*[]ShareResultQuery, error)
	GetById(shareId *uuid.UUID) (*[]ShareResultQuery, error)
}
type ShareResultQuery struct {
	SpendingTypeId uuid.UUID
	NameTh         string
	NameEn         string
}

type ShareInsertDB struct {
	NameTh         string
	NameEn         string
	SpendingTypeId uuid.UUID
}
