package routes

import (
	"money-service/controllers"
	"money-service/middlewares"

	"github.com/gofiber/fiber/v2"
)

func IncomeRoute(app fiber.Router, controllers controllers.IIncomeController) {
	api := app.Group("/incomes")
	api.Use(middlewares.JWTMiddleware())
	api.Get("/", controllers.GetIncomesByUser())
}
