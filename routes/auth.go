package routes

// route login
// route register // create user
// route logout // backlist token

import (
	"money-service/controllers"
	"money-service/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app fiber.Router, controllers controllers.IAuthController) {
	api := app.Group("/auth")
	api.Use(middlewares.JWTMiddleware())
	api.Post("/login", controllers.Login())
	api.Post("/register", controllers.Register())
	api.Post("/logout", controllers.Logout())
}
