package controllers

import (
	"fmt"
	"money-service/interfaces"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
)

type IAuthController struct {
	service *services.IAuthService
}

func AuthController(service *services.IAuthService) IAuthController {
	return IAuthController{service}
}
func (s IAuthController) Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := interfaces.AuthRegisterRequest{}
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		fmt.Println(payload)

		res, err := s.service.Register(interfaces.AuthRegisterInsert{
			UserName:    payload.UserName,
			Email:       payload.Email,
			LastName:    payload.LastName,
			FirstName:   payload.FirstName,
			DisplayName: payload.DisplayName,
			Password:    payload.Password,
			Profile:     payload.Profile,
		})
		fmt.Print(res)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(res)
	}
}
func (s IAuthController) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := interfaces.AuthLoginRequest{}
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		res, err := s.service.Login(payload.Credential, payload.Password)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(res)
	}
}
func (s IAuthController) Logout() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
