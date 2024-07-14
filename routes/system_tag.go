package routes

import (
	"fmt"
	"money-service/controllers"

	"github.com/gofiber/fiber/v2"
)

func SystemTagRoute(app fiber.Router, controllers controllers.ISystemTagController) {
	fmt.Print("route ")
	api := app.Group("/users")
	api.Get("/:id", controllers.GetSystemTagById())
}
