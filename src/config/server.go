package config

import (
	"fmt"
	"log"
	_ "money-service/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var app = fiber.New()

func GetInstanceServer() (fiber.Router, *Config) {
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000,http://localhost:3000/", // ระบุ origin ที่ต้องการ
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, X-API-KEY,Authorization",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	}))
	api := app.Group(fmt.Sprintf("/api/%s", cfg.API_VERSION))
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("welcome to my api")
	})

	return api, cfg
}

// @title Fiber Swagger API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /api/v1

func StartServer() {

	app.Get("/swagger/*", swagger.HandlerDefault)

	if err := app.Listen(":" + cfg.SERVER_PORT); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
