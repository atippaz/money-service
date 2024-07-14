package controllers

import (
	"money-service/services"

	"github.com/gofiber/fiber/v2"
)

type IExpenseController interface {
	GetExpensesByUser() fiber.Handler
}

type expenseController struct {
	service *services.IExpenseService
}

func ExpenseController(service *services.IExpenseService) IExpenseController {
	return expenseController{service}
}
func (s expenseController) GetExpensesByUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := ""
		res, err := s.service.GetExpensesByUser(userId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
