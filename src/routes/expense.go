package routes

import (
	controllers "money-service/src/controllers/expense"
	"money-service/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ExpenseRoute(app fiber.Router, controllers controllers.FiberExpenseController, jwtMiddleware middlewares.JWTMiddleware) {
	api := app.Group("/expenses")
	api.Use(jwtMiddleware.MiddleWare())
	api.Get("/", controllers.GetExpensesByUser())
	api.Post("/", controllers.CreateExpense())
	api.Patch("/", controllers.CreateExpense())

	// update in 30 minute
	// delete in 1 day
}
