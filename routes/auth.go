package routes

// route login
// route register // create user
// route logout // backlist token

import (
	"money-service/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app fiber.Router, controllers controllers.IAuthController) {
	api := app.Group("/auth")
	api.Post("/login", controllers.Login())
	api.Post("/register", controllers.Register())
	api.Post("/logout", controllers.Logout())
}
