package controllers

import (
	"money-service/interfaces"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
		res, err := s.service.GetExpensesByUser(uuid.MustParse(userId))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}

func (s IExpenseController) CreateExpense() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var payload interfaces.ExpenseInsertRequest
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		userId := "1"
		res, err := s.service.CreateExpense(uuid.MustParse(userId), interfaces.ExpenseInsertDb{
			TagId: payload.TagId,
			Value: payload.Value,
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
