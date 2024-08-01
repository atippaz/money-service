package services

import (
	repositories "money-service/src/repositories/income"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type IncomeService interface {
	CreateIncome(id uuid.UUID, payload IncomeInsert) (*uuid.UUID, error)
	GetIncomesByUser(id uuid.UUID, startDate *time.Time, endDate *time.Time) (*[]IncomeResult, error)
	GetSummary(id uuid.UUID, startDate *time.Time, endDate *time.Time) (*[]IncomeSummaryResult, error)
}

type incomeService struct {
	repo repositories.IncomeRepository
}
type IncomeResult struct {
	IncomeId    uuid.UUID
	CreatedDate time.Time
	TagId       uuid.UUID
	Value       decimal.Decimal
	UserOwner   uuid.UUID
}
type IncomeSummaryResult struct {
	TagId uuid.UUID
	Value decimal.Decimal
}
type IncomeInsert struct {
	TagId     uuid.UUID
	Value     decimal.Decimal
	UserOwner uuid.UUID
}
