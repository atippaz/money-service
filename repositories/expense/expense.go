package repositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ExpenseRepository interface {
	GetExpensesByUser(userId uuid.UUID) (*[]ExpenseResultQuery, error)
	CreateExpense(userId uuid.UUID, payload ExpenseInsertDb) (*uuid.UUID, error)
}

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
