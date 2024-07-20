package controllers

import (
	incomeService "money-service/src/services/income"

	"github.com/gofiber/fiber/v2"
)

type IncomeController[T any] interface {
	CreateIncome() T
	GetIncomesByUser() T
}

type incomeController struct {
	service incomeService.IncomeService
}

type FiberIncomeController interface {
	IncomeController[fiber.Handler]
}
