package interfaces

import "github.com/google/uuid"

type SpendingTypeResultResponse struct {
	SpendingTypeId uuid.UUID `json:"spendingTypeId"`
	NameTh         string    `json:"nameTh"`
	NameEn         string    `json:"nameEng"`
}
