package controllers

import (
	"fmt"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
)

type ISystemTagController interface {
	GetSystemTagById() fiber.Handler
}

type systemTagController struct {
	service *services.ISystemTagService
}

func SystemTagController(service *services.ISystemTagService) ISystemTagController {
	return systemTagController{service}
}
func (s systemTagController) GetSystemTagById() fiber.Handler {
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
