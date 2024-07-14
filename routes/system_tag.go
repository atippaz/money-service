package routes

import (
	"money-service/controllers"

	"github.com/gofiber/fiber/v2"
)

func SystemTagRoute(app fiber.Router, controllers controllers.ISystemTagController) {
	api := app.Group("/system_tags")
	api.Get("/", controllers.GetAllSystemTags())
}
