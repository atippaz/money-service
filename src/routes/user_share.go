package routes

import (
	controllers "money-service/src/controllers/user_share"
	"money-service/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func UserShareRoute(app fiber.Router, controllers controllers.FiberUserShareController, jwtMiddleware middlewares.JWTMiddleware) {
	api := app.Group("/user-shares")
	api.Use(jwtMiddleware.MiddleWare())
	api.Get("/", controllers.GetAllUserShareHandler())
	api.Post("/", controllers.InsertUserShareHandler())
}
