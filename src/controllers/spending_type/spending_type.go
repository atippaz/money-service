package controllers

import (
	spendingType "money-service/src/services/spending_type"

	"github.com/gofiber/fiber/v2"
)

func NewFiberSpendingTypeController(service spendingType.SpendingTypeService) FiberSpendingTypeController {
	return &spendingTypeController{service}
}
func (s spendingTypeController) GetSpendingHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		res, err := s.service.GetSpendingTypes()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
