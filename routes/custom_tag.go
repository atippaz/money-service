package routes

import (
	"fmt"
	"money-service/controllers"

	"github.com/gofiber/fiber/v2"
)

func CustomTagRoute(app fiber.Router, controllers controllers.ICustomTagController) {
	fmt.Print("route ")
	api := app.Group("/users")
	api.Get("/:id", controllers.GetCustomTagById())
}
