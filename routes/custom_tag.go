package routes

import (
	"money-service/controllers"
	"money-service/middlewares"

	"github.com/gofiber/fiber/v2"
)

func CustomTagRoute(app fiber.Router, controllers controllers.ICustomTagController) {
	api := app.Group("/custom_tags")
	api.Use(middlewares.JWTMiddleware())
	api.Get("/:id", controllers.GetCustomTagsByUser())
}
