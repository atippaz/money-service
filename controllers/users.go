package controllers

import (
	"fmt"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
)

type IUserController interface {
	GetUserById() fiber.Handler
}

type userController struct {
	service *services.IUserService
}

func UserController(service *services.IUserService) IUserController {
	return userController{service}
}
func (s userController) GetUserById() fiber.Handler {
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
