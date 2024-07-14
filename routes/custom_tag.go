package routes

import (
	"money-service/controllers"

	"github.com/gofiber/fiber/v2"
)

func CustomTagRoute(app fiber.Router, controllers controllers.ICustomTagController) {
	api := app.Group("/custom_tags")
	api.Get("/:id", controllers.GetCustomTagById())
}
