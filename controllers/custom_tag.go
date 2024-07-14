package controllers

import (
	"fmt"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
)

type ICustomTagController interface {
	GetCustomTagById() fiber.Handler
}

type customTagController struct {
	service *services.ICustomTagService
}

func CustomTagController(service *services.ICustomTagService) ICustomTagController {
	return customTagController{service}
}
func (s customTagController) GetCustomTagById() fiber.Handler {
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
