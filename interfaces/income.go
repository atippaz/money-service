package interfaces

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type IncomeResultQuery struct {
	IncomeId    uuid.UUID
	CreatedDate time.Time
	TagId       uuid.UUID
	Value       decimal.Decimal
	UserOwner   uuid.UUID
}

type IncomeInsertDb struct {
	TagId     uuid.UUID
	Value     decimal.Decimal
	UserOwner uuid.UUID
}

type IncomeInsertRequest struct {
	TagId     uuid.UUID       `json:"tagId"`
	Value     decimal.Decimal `json:"value"`
	UserOwner uuid.UUID       `json:"userOwner"`
}
