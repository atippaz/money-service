package routes

import (
	"money-service/controllers"
	"money-service/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ExpenseRoute(app fiber.Router, controllers controllers.IExpenseController) {
	api := app.Group("/expenses")
	api.Use(middlewares.JWTMiddleware())
	api.Get("/", controllers.GetExpensesByUser())
	api.Post("/", controllers.GetExpensesByUser())
	// delete in 1 day
	// update in 30 minute
}
