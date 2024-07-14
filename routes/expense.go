package routes

import (
	"fmt"
	"money-service/controllers"

	"github.com/gofiber/fiber/v2"
)

func ExpenseRoute(app fiber.Router, controllers controllers.IExpenseController) {
	fmt.Print("route ")
	api := app.Group("/users")
	api.Get("/:id", controllers.GetUserById())
}
