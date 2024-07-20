package routes

import (
	"fmt"
	controllers "money-service/controllers/user"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app fiber.Router, controllers controllers.FiberUserController) {
	fmt.Print("route ")
	api := app.Group("/users")
	// api.Use(middlewares.JWTMiddleware())
	api.Get("/profile", controllers.GetUserById())
	api.Delete("/", controllers.DeActiveAccount())
	// update
}
