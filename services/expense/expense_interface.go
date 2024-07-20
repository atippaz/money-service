package services

import (
	repositories "money-service/repositories/expense"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ExpenseService interface {
	GetExpensesByUser(userId uuid.UUID) (*[]ExpenseResult, error)
	CreateExpense(userId uuid.UUID, payload ExpenseInsert) (*uuid.UUID, error)
}

type expenseService struct {
	repo repositories.ExpenseRepository
}

type ExpenseResult struct {
	ExpenseId   uuid.UUID
	CreatedDate time.Time
	TagId       uuid.UUID
	Value       decimal.Decimal
	UserOwner   uuid.UUID
}
type ExpenseInsert struct {
	TagId uuid.UUID
	Value decimal.Decimal
}
