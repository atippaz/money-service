package controllers

import (
	"fmt"
	"money-service/interfaces"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type IIncomeController struct {
	service *services.IIncomeService
}

func IncomeController(service *services.IIncomeService) IIncomeController {
	return IIncomeController{service}
}
func (s IIncomeController) CreateIncome() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//todo get uuid from user middleware
		//todo get payload form body
		payload := interfaces.IncomeInsertDb{}
		id := uuid.New()
		res, err := s.service.CreateIncome(id, payload)
		fmt.Print(res)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
func (s IIncomeController) GetIncomesByUser() fiber.Handler {
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
