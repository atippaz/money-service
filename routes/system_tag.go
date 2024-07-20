package routes

import (
	"money-service/controllers"

	"github.com/gofiber/fiber/v2"
)

func SystemTagRoute(app fiber.Router, controllers controllers.FiberSystemTagController) {
	api := app.Group("/system-tags")
	api.Get("/", controllers.GetAllSystemTags())
}
