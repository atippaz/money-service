package interfaces

import "github.com/google/uuid"

type SpendingTypeResult struct {
	SpendingTypeId uuid.UUID
	NameTh         string
	NameEn         string
}
