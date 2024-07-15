package controllers

import (
	"fmt"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type IUserController struct {
	service *services.IUserService
}

func UserController(service *services.IUserService) IUserController {
	return IUserController{service}
}
func (s IUserController) GetUserById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Locals("user").(uuid.UUID)
		res, err := s.service.GetUserById(id)
		fmt.Print(res)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
func (s IUserController) DeActiveAccount() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// todo get id form middle ware
		id := ""
		res, err := s.service.DeActiveAccount(id)
		fmt.Print(res)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
