package controllers

import (
	expenseSevice "money-service/src/services/expense"
	jwt "money-service/src/utils/jwt"

	"github.com/gofiber/fiber/v2"
)

func NewFiberExpenseController(service expenseSevice.ExpenseService) FiberExpenseController {
	return &expenseController{service}
}

func (s expenseController) GetExpensesByUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals("user").(*jwt.AuthClaims)
		res, err := s.service.GetExpensesByUser(claims.UserId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		results := []ExpenseResult{}
		for _, result := range *res {
			results = append(results, ExpenseResult{
				CreatedDate: result.CreatedDate,
				ExpenseId:   result.ExpenseId.String(),
				TagId:       result.TagId.String(),
				UserOwner:   result.UserOwner.String(),
				Value:       result.Value,
			})
		}
		return c.JSON(results)
	}
}

func (s expenseController) CreateExpense() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var payload ExpenseInsertRequest
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		claims := c.Locals("user").(*jwt.AuthClaims)
		res, err := s.service.CreateExpense(claims.UserId, expenseSevice.ExpenseInsert{
			TagId: payload.TagId,
			Value: payload.Value,
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
