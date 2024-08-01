package controllers

import (
	incomeService "money-service/src/services/income"
	Datetime "money-service/src/utils/datetime"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type IncomeController[T any] interface {
	CreateIncome() T
	GetIncomesByUser() T
	GetSummary() T
}

type incomeController struct {
	service  incomeService.IncomeService
	datetime Datetime.Datetime
}

type FiberIncomeController interface {
	IncomeController[fiber.Handler]
}

type IncomeResult struct {
	CreatedDate time.Time       `json:"createdDate" `
	IncomeId    string          `json:"incomeId" `
	TagId       string          `json:"tagId" `
	UserOwner   string          `json:"userOwner" `
	Value       decimal.Decimal `json:"value"`
}

type IncomeSummaryResult struct {
	TagId          string          `json:"tagId" `
	Value          decimal.Decimal `json:"value"`
	SpendingTypeId string          `json:"spendingTypeId"`
}
type IncomeInsertRequest struct {
	TagId uuid.UUID       `json:"tagId"`
	Value decimal.Decimal `json:"value"`
	Date  string          `json:"date"`
}
