package routes

import (
	"money-service/controllers"

	"github.com/gofiber/fiber/v2"
)

func SpendingTypeController(app fiber.Router, controllers controllers.ISpendingTypeContoller) {
	api := app.Group("/spending-types")
	api.Get("/", controllers.GetSpendingHandler())
}
