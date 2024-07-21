package controllers

import (
	userService "money-service/src/services/user"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func NewFiberUserController(service userService.UserService) FiberUserController {
	return &userController{service}
}
func (s userController) GetUserById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Locals("user").(uuid.UUID)
		res, err := s.service.GetUserById(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
func (s userController) DeActiveAccount() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// todo get id form middle ware
		id := ""
		res, err := s.service.DeActiveAccount(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
