package controllers

import (
	"money-service/services"

	"github.com/gofiber/fiber/v2"
)

type SystemTagController[T any] interface {
	GetAllSystemTags() T
}
type systemTagController struct {
	service *services.SystemTagService
}

// implement
type FiberSystemTagController interface {
	SystemTagController[fiber.Handler]
}

func NewFiberSystemTagController(service *services.SystemTagService) FiberSystemTagController {
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
