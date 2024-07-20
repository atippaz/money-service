package interfaces

import "github.com/google/uuid"

type SpendingTypeResultQuery struct {
	SpendingTypeId uuid.UUID
	NameTh         string
	NameEn         string
}
type SpendingTypeResultResponse struct {
	SpendingTypeId uuid.UUID `json:"spendingTypeId"`
	NameTh         string    `json:"nameTh"`
	NameEn         string    `json:"nameEng"`
}
