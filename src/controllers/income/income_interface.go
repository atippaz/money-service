package controllers

import (
	incomeService "money-service/src/services/income"
	Datetime "money-service/src/utils/datetime"

	"github.com/gofiber/fiber/v2"
)

type IncomeController[T any] interface {
	CreateIncome() T
	GetIncomesByUser() T
}

type incomeController struct {
	service  incomeService.IncomeService
	datetime Datetime.Datetime
}

type FiberIncomeController interface {
	IncomeController[fiber.Handler]
}
