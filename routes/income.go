package routes

import (
	"money-service/controllers"

	"github.com/gofiber/fiber/v2"
)

func IncomeRoute(app fiber.Router, controllers controllers.IIncomeController) {
	api := app.Group("/incomes")
	api.Get("/", controllers.GetIncomesByUser())
}
