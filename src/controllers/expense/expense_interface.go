package controllers

import (
	expenseSevice "money-service/src/services/expense"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ExpenseController[T fiber.Handler] interface {
	GetExpensesByUser() T
	CreateExpense() T
}

type expenseController struct {
	service expenseSevice.ExpenseService
}

type FiberExpenseController interface {
	ExpenseController[fiber.Handler]
}
type ExpenseInsertRequest struct {
	TagId uuid.UUID       `json:"tagId"`
	Value decimal.Decimal `json:"value"`
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
