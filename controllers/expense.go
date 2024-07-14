package controllers

import (
	"money-service/interfaces"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
)

type IExpenseController struct {
	service *services.IExpenseService
}

func ExpenseController(service *services.IExpenseService) IExpenseController {
	return IExpenseController{service}
}
func (s IExpenseController) GetExpensesByUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := ""
		res, err := s.service.GetExpensesByUser(userId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
func (s IExpenseController) CreateExpense() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := ""
		res, err := s.service.CreateExpense(userId, interfaces.ExpenseResultInsertDb{})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
