package interfaces

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ExpenseResultQuery struct {
	ExpenseId   uuid.UUID
	CreatedDate time.Time
	TagId       uuid.UUID
	Value       decimal.Decimal
	UserOwner   uuid.UUID
}

type ExpenseInsertDb struct {
	TagId uuid.UUID
	Value decimal.Decimal
}

type ExpenseInsertRequest struct {
	TagId uuid.UUID       `json:"tagId"`
	Value decimal.Decimal `json:"value"`
}
