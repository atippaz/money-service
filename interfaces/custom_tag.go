package interfaces

import "github.com/google/uuid"

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
