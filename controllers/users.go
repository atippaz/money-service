package controllers

import (
	"fmt"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserController[T any] interface {
	GetUserById() T
	DeActiveAccount() T
}
type userController struct {
	service *services.UserService
}
type FiberUserController interface {
	UserController[fiber.Handler]
}

func NewFiberUserController(service *services.UserService) FiberUserController {
	return &userController{service}
}
func (s userController) GetUserById() fiber.Handler {
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
func (s userController) DeActiveAccount() fiber.Handler {
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
