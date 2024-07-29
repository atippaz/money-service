package repositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type IncomeRepository interface {
	GetIncomesByUser(userId uuid.UUID, startDate *time.Time, endDate *time.Time) (*[]IncomeResultQuery, error)
	CreateIncome(id uuid.UUID, payload IncomeInsertDb) (*uuid.UUID, error)
}

type IncomeResultQuery struct {
	IncomeId    uuid.UUID
	CreatedDate time.Time
	TagId       uuid.UUID
	Value       decimal.Decimal
	UserOwner   uuid.UUID
}

type IncomeInsertDb struct {
	TagId uuid.UUID
	Value decimal.Decimal
}
