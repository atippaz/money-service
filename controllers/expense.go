package controllers

import (
	"fmt"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
)

type IExpenseController interface {
	GetUserById() fiber.Handler
}

type expenseController struct {
	service *services.IExpenseService
}

func ExpenseController(service *services.IExpenseService) IExpenseController {
	return expenseController{service}
}
func (s expenseController) GetUserById() fiber.Handler {
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
