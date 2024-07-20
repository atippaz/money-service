package routes

import (
	controllers "money-service/src/controllers/custom_tag"
	"money-service/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func CustomTagRoute(app fiber.Router, controllers controllers.FiberCustomTagController, jwtMiddleware middlewares.JWTMiddleware) {
	api := app.Group("/custom-tags")
	api.Use(jwtMiddleware.MiddleWare())
	api.Get("/", controllers.GetCustomTagsByUser())
	api.Post("/", controllers.CreateCustomTag())

	// todo
	// api.Patch("/:custom_tag_id")

	// update by id
	// delete by id todo soft delete isdeactive status
}
