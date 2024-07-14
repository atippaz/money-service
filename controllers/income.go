package controllers

import (
	"fmt"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
)

type IIncomeController interface {
	GetIncomesByUser() fiber.Handler
}

type incomeController struct {
	service *services.IIncomeService
}

func IncomeController(service *services.IIncomeService) IIncomeController {
	return incomeController{service}
}
func (s incomeController) GetIncomesByUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := ""
		res, err := s.service.GetIncomesByUser(id)
		fmt.Print(res)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
