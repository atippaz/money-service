package controllers

import (
	incomeService "money-service/src/services/income"
	jwtUtils "money-service/src/utils/jwt"

	"github.com/gofiber/fiber/v2"
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
		// if err != nil {
		// 	return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		// }
		// return c.JSON(res)
		return c.JSON("")
	}
}
func (s incomeController) GetIncomesByUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals("user").(*jwtUtils.AuthClaims)
		res, err := s.service.GetIncomesByUser(claims.UserId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
