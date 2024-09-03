package routes

import (
	controllers "money-service/src/controllers/line"

	"github.com/gofiber/fiber/v2"
)

func LineRoute(app fiber.Router, controllers controllers.FiberLineTypeController) {
	app.Post("/webhook", controllers.Webhook())

	// delete in 1 day
	// update in 30 minute
}
