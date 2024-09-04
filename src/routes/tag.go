package routes

import (
	controllers "money-service/src/controllers/tag"
	"money-service/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func CustomTagRoute(app fiber.Router, controllers controllers.FiberTagController, jwtMiddleware middlewares.JWTMiddleware) {
	api := app.Group("/tags")
	api.Use(jwtMiddleware.MiddleWare())
	api.Get("/", controllers.GetTagsByUser())
	api.Post("/", controllers.CreateTag())

	// todo
	// api.Patch("/:custom_tag_id")

	// update by id
	// delete by id todo soft delete isdeactive status
}
