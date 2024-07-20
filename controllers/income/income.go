package controllers

import (
	"fmt"
	incomeService "money-service/services/income"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func NewFiberIncomeController(service incomeService.IncomeService) FiberIncomeController {
	return &incomeController{service}
}
func (s incomeController) CreateIncome() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// //todo get uuid from user middleware
		// //todo get payload form body
		// payload := IncomeInsertDb{}
		// id := uuid.New()
		// res, err := s.service.CreateIncome(id, payload)
		// fmt.Print(res)
		// if err != nil {
		// 	return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		// }
		// return c.JSON(res)
		return c.JSON("")
	}
}
func (s incomeController) GetIncomesByUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// get uuid from user middleware
		id := uuid.New()
		res, err := s.service.GetIncomesByUser(id)
		fmt.Print(res)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
