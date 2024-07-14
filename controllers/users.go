package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func UserController(app fiber.Router) {
	app.Get("/users", func(c *fiber.Ctx) error {
		return c.SendString("List of users")
	})
}
