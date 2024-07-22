package controllers

import (
	authSevice "money-service/src/services/auth"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func NewFiberAuthController(service authSevice.AuthService, validate *validator.Validate) FiberAuthController {
	return &authController{service, validate}
}

func (s authController) Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := AuthRegisterRequest{}
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		if err := s.validate.Struct(payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		res, err := s.service.Register(authSevice.AuthRegisterInsert{
			UserName:    payload.UserName,
			Email:       payload.Email,
			LastName:    payload.LastName,
			FirstName:   payload.FirstName,
			DisplayName: payload.DisplayName,
			Password:    payload.Password,
			Profile:     payload.Profile,
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(res)
	}
}
func (s authController) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := AuthLoginRequest{}
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		if err := s.validate.Struct(payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		res, err := s.service.Login(payload.Credential, payload.Password)
		if err != nil {
			return c.Status(400).JSON("not found")
		}
		return c.JSON(res)
	}
}
func (s authController) Logout() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
