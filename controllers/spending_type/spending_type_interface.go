package controllers

import (
	spendingType "money-service/services/spending_type"

	"github.com/gofiber/fiber/v2"
)

type SpendingTypeController[T any] interface {
	GetSpendingHandler() T
}

type spendingTypeController struct {
	service spendingType.SpendingTypeService
}

type FiberSpendingTypeController interface {
	SpendingTypeController[fiber.Handler]
}
