package controllers

import (
	"fmt"
	customTagSevice "money-service/src/services/custom_tag"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func NewFiberCustomTagController(service customTagSevice.CustomTagService) FiberCustomTagController {
	return &customTagController{service}
}
func (s customTagController) GetCustomTagsByUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// user := c.Locals("user").(string)
		// fmt.Println(user)
		id := c.Params("id")
		res, err := s.service.GetCustomTagsByUser(id)
		fmt.Print(res)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
func (s customTagController) CreateCustomTag() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// user := c.Locals("user").(string)
		id := uuid.MustParse("e8bc8014-4a5a-4e91-a6d2-3b5b6f77d3e0")
		res, err := s.service.CreateCustomTag(id)
		fmt.Print(res)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
