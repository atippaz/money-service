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
