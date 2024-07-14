package interfaces

import "github.com/google/uuid"

type SpendingTypeResultQuery struct {
	SpendingTypeId uuid.UUID
	NameTh         string
	NameEn         string
}
