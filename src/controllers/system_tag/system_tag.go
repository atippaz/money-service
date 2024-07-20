package controllers

import (
	SystemTagService "money-service/src/services/system_tag"

	"github.com/gofiber/fiber/v2"
)

func NewFiberSystemTagController(service SystemTagService.SystemTagService) FiberSystemTagController {
	return &systemTagController{service}
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
