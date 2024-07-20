package interfaces

import "github.com/google/uuid"

type CustomTagInsertRequest struct {
	NameTh         string
	NameEn         string
	SpendingTypeId uuid.UUID
}
