package controllers

import (
	"fmt"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
)

type IUserContoller interface {
	GetUserById() fiber.Handler
}

type userContoller struct {
	service *services.IUserService
}

func UserContoller(service *services.IUserService) IUserContoller {
	return userContoller{service}
}
func (s userContoller) GetUserById() fiber.Handler {
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
