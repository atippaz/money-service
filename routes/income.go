package routes

import (
	"fmt"
	"money-service/controllers"

	"github.com/gofiber/fiber/v2"
)

func IncomeRoute(app fiber.Router, controllers controllers.IIncomeController) {
	fmt.Print("route ")
	api := app.Group("/users")
	api.Get("/:id", controllers.GetIncomeById())
}
