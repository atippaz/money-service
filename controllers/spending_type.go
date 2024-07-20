package controllers

import (
	"fmt"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
)

type SpendingTypeController[T any] interface {
	GetSpendingHandler() T
}

type spendingTypeController struct {
	service *services.SpendingTypeService
}

// implement
type FiberSpendingTypeController interface {
	SpendingTypeController[fiber.Handler]
}

func NewFiberSpendingTypeController(service *services.SpendingTypeService) FiberSpendingTypeController {
	return &spendingTypeController{service}
}
func (s spendingTypeController) GetSpendingHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		res, err := s.service.GetSpendingTypes()
		fmt.Print(res)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
