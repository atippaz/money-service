package controllers

import (
	"fmt"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
)

type IIncomeController interface {
	GetIncomeById() fiber.Handler
}

type incomeController struct {
	service *services.IIncomeService
}

func IncomeController(service *services.IIncomeService) IIncomeController {
	return incomeController{service}
}
func (s incomeController) GetIncomeById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		res, err := s.service.GetUserById(id)
		fmt.Print(res)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
