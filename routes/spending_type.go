package routes

import (
	"money-service/controllers"

	"github.com/gofiber/fiber/v2"
)

func SpendingTypeRoute(app fiber.Router, controllers controllers.ISpendingTypeController) {
	api := app.Group("/spending-types")
	api.Get("/", controllers.GetSpendingHandler())
}
