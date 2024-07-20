package config

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var app = fiber.New()

func GetInstanceServer() (fiber.Router, *Config) {
	app.Use(logger.New())
	app.Use(recover.New())

	api := app.Group(fmt.Sprintf("/api/%s", cfg.API_VERSION))
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("welcome to my api")
	})

	return api, cfg
}
func StartServer() {
	if err := app.Listen(":" + cfg.SERVER_PORT); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
