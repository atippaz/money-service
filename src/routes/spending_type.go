package routes

import (
	controllers "money-service/src/controllers/spending_type"

	"github.com/gofiber/fiber/v2"
)

func SpendingTypeRoute(app fiber.Router, controllers controllers.FiberSpendingTypeController) {
	api := app.Group("/spending-types")
	api.Get("/", controllers.GetSpendingHandler())
}
