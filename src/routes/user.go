package routes

import (
	controllers "money-service/src/controllers/user"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app fiber.Router, controllers controllers.FiberUserController) {
	api := app.Group("/users")
	// api.Use(middlewares.JWTMiddleware())
	api.Get("/profile", controllers.GetUserById())
	api.Delete("/deactivate", controllers.DeActiveAccount())
	// update
}
