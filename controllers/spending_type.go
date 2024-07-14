package controllers

import (
	"fmt"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
)

type ISpendingTypeContoller interface {
	GetSpendingHandler() fiber.Handler
}

type spendingTypeContoller struct {
	service *services.ISpendingTypeService
}

func SpendingTypeContoller(service *services.ISpendingTypeService) ISpendingTypeContoller {
	return spendingTypeContoller{service}
}
func (s spendingTypeContoller) GetSpendingHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		res, err := s.service.GetSpendingTypes()
		fmt.Print(res)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
