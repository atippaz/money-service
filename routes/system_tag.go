package routes

import (
	controllers "money-service/controllers/system_tag"

	"github.com/gofiber/fiber/v2"
)

func SystemTagRoute(app fiber.Router, controllers controllers.FiberSystemTagController) {
	api := app.Group("/system-tags")
	api.Get("/", controllers.GetAllSystemTags())
}
