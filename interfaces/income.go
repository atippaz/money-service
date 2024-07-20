package interfaces

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type IncomeInsertRequest struct {
	TagId     uuid.UUID       `json:"tagId"`
	Value     decimal.Decimal `json:"value"`
	UserOwner uuid.UUID       `json:"userOwner"`
}
