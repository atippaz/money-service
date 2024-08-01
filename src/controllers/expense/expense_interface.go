package controllers

import (
	expenseSevice "money-service/src/services/expense"
	Datetime "money-service/src/utils/datetime"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ExpenseController[T fiber.Handler] interface {
	GetExpensesByUser() T
	CreateExpense() T
	GetSummary() T
}

type expenseController struct {
	service  expenseSevice.ExpenseService
	datetime Datetime.Datetime
}

type FiberExpenseController interface {
	ExpenseController[fiber.Handler]
}
type ExpenseInsertRequest struct {
	TagId uuid.UUID       `json:"tagId"`
	Value decimal.Decimal `json:"value"`
	Date  string          `json:"date"`
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
type ExpenseResult struct {
	CreatedDate time.Time       `json:"createdDate" `
	ExpenseId   string          `json:"ำxpenseId" `
	TagId       string          `json:"tagId" `
	UserOwner   string          `json:"userOwner" `
	Value       decimal.Decimal `json:"value"`
}
type ExpenseSummaryResult struct {
	TagId          string          `json:"tagId" `
	Value          decimal.Decimal `json:"value"`
	SpendingTypeId string          `json:"spendingTypeId"`
}
