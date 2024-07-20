package routes

import (
	"money-service/controllers"

	"github.com/gofiber/fiber/v2"
)

func IncomeRoute(app fiber.Router, controllers controllers.FiberIncomeController) {
	api := app.Group("/incomes")
	// api.Use(middlewares.JWTMiddleware())
	api.Get("/", controllers.GetIncomesByUser())
	api.Post("/", controllers.CreateIncome())
	// delete in 1 day
	// update in 30 minute
}
