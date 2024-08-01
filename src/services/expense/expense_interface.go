package services

import (
	repositories "money-service/src/repositories/expense"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ExpenseService interface {
	GetExpensesByUser(userId uuid.UUID, startDate *time.Time, endDate *time.Time) (*[]ExpenseResult, error)
	CreateExpense(userId uuid.UUID, payload ExpenseInsert) (*uuid.UUID, error)
	GetSummary(id uuid.UUID, startDate *time.Time, endDate *time.Time) (*[]ExpenseSummaryResult, error)
}

type expenseService struct {
	repo repositories.ExpenseRepository
}
type ExpenseSummaryResult struct {
	TagId uuid.UUID
	Value decimal.Decimal
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
