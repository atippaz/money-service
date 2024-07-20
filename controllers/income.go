package controllers

import (
	"fmt"
	"money-service/interfaces"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type IncomeController[T any] interface {
	CreateIncome() T
	GetIncomesByUser() T
}

type incomeController struct {
	service *services.IncomeService
}

// implement
type FiberIncomeController interface {
	IncomeController[fiber.Handler]
}

func NewFiberIncomeController(service *services.IncomeService) FiberIncomeController {
	return &incomeController{service}
}
func (s incomeController) CreateIncome() fiber.Handler {
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
