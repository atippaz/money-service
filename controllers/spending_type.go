package controllers

import (
	"fmt"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
)

type ISpendingTypeController interface {
	GetSpendingHandler() fiber.Handler
}

type spendingTypeController struct {
	service *services.ISpendingTypeService
}

func SpendingTypeController(service *services.ISpendingTypeService) ISpendingTypeController {
	return spendingTypeController{service}
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
