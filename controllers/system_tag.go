package controllers

import (
	"money-service/services"

	"github.com/gofiber/fiber/v2"
)

type ISystemTagController struct {
	service *services.ISystemTagService
}

func SystemTagController(service *services.ISystemTagService) ISystemTagController {
	return ISystemTagController{service}
}
func (s ISystemTagController) GetAllSystemTags() fiber.Handler {
	return func(c *fiber.Ctx) error {
		res, err := s.service.GetAllSystemTags()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
