package routes

import (
	controllers "money-service/src/controllers/user"
	"money-service/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app fiber.Router, controllers controllers.FiberUserController, jwtMiddleware middlewares.JWTMiddleware) {
	api := app.Group("/users")
	api.Use(jwtMiddleware.MiddleWare())
	api.Get("/profile", controllers.GetUserById())
	api.Delete("/deactivate", controllers.DeActiveAccount())
	// update
}
