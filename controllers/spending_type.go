package controllers

import (
	"fmt"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
)

type ISpendingTypeController struct {
	service *services.ISpendingTypeService
}

func SpendingTypeController(service *services.ISpendingTypeService) ISpendingTypeController {
	return ISpendingTypeController{service}
}
func (s ISpendingTypeController) GetSpendingHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		res, err := s.service.GetSpendingTypes()
		fmt.Print(res)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
