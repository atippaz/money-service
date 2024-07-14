package routes

import (
	"money-service/controllers"

	"github.com/gofiber/fiber/v2"
)

func ExpenseRoute(app fiber.Router, controllers controllers.IExpenseController) {
	api := app.Group("/expenses")
	api.Get("/", controllers.GetExpensesByUser())
}
