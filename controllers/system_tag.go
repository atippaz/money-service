package controllers

import (
	"money-service/services"

	"github.com/gofiber/fiber/v2"
)

type ISystemTagController interface {
	GetAllSystemTags() fiber.Handler
}

type systemTagController struct {
	service *services.ISystemTagService
}

func SystemTagController(service *services.ISystemTagService) ISystemTagController {
	return systemTagController{service}
}
func (s systemTagController) GetAllSystemTags() fiber.Handler {
	return func(c *fiber.Ctx) error {
		res, err := s.service.GetAllSystemTags()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
