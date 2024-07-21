package controllers

import (
	customTagSevice "money-service/src/services/custom_tag"
	jwtUtils "money-service/src/utils/jwt"

	"github.com/gofiber/fiber/v2"
)

func NewFiberCustomTagController(service customTagSevice.CustomTagService) FiberCustomTagController {
	return &customTagController{service}
}
func (s customTagController) GetCustomTagsByUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals("user").(*jwtUtils.AuthClaims)
		res, err := s.service.GetCustomTagsByUser(claims.UserId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
func (s customTagController) CreateCustomTag() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals("user").(*jwtUtils.AuthClaims)
		res, err := s.service.CreateCustomTag(claims.UserId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
