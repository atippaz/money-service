package controllers

import (
	userService "money-service/src/services/user"
	jwtUtils "money-service/src/utils/jwt"

	"github.com/gofiber/fiber/v2"
)

func NewFiberUserController(service userService.UserService) FiberUserController {
	return &userController{service}
}
func (s userController) GetUserById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals("user").(*jwtUtils.AuthClaims)
		res, err := s.service.GetUserById(claims.UserId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
func (s userController) DeActiveAccount() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals("user").(*jwtUtils.AuthClaims)
		res, err := s.service.DeActiveAccount(claims.UserId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
