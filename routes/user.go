package routes

import (
	"money-service/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app fiber.Router, controllers controllers.IUserContoller) {
	api := app.Group("/users")
	api.Get("/:id", controllers.GetUserById(""))
}
