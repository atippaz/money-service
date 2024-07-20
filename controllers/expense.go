package controllers

import (
	"money-service/interfaces"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ExpenseController[T fiber.Handler] interface {
	GetExpensesByUser() T
	CreateExpense() T
}

type expenseController struct {
	service *services.ExpenseService
}

// implement
type FiberExpenseController interface {
	ExpenseController[fiber.Handler]
}

func NewFiberExpenseController(service *services.ExpenseService) FiberExpenseController {
	return &expenseController{service}
}

func (s expenseController) GetExpensesByUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals("user").(*interfaces.AuthClaims)
		res, err := s.service.GetExpensesByUser(claims.UserId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}

func (s expenseController) CreateExpense() fiber.Handler {
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
