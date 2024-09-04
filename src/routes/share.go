package routes

import (
	controllers "money-service/src/controllers/share"
	"money-service/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ShareRoute(app fiber.Router, controllers controllers.FiberShareController, jwtMiddleware middlewares.JWTMiddleware) {
	api := app.Group("/shares")
	api.Use(jwtMiddleware.MiddleWare())
	api.Get("/:id", controllers.GetByIdShareHandler())
	api.Post("/", controllers.InsertShareHandler())
	api.Get("/", controllers.GetAllShareHandler())
}
