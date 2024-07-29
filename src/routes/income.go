package routes

import (
	controllers "money-service/src/controllers/income"
	"money-service/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func IncomeRoute(app fiber.Router, controllers controllers.FiberIncomeController, jwtMiddleware middlewares.JWTMiddleware) {
	api := app.Group("/incomes")
	api.Use(jwtMiddleware.MiddleWare())
	api.Get("/", controllers.GetIncomesByUser())
	api.Post("/", controllers.CreateIncome())
	// delete in 1 day
	// update in 30 minute
}
