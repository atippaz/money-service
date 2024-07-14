package routes

import (
	"fmt"
	"money-service/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app fiber.Router, controllers controllers.IUserController) {
	fmt.Print("route ")
	api := app.Group("/users")
	api.Get("/:id", controllers.GetUserById())
}