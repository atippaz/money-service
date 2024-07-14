package initials

import (
	"money-service/config"
	"money-service/controllers"
	repositories "money-service/repositories/spending_type"
	"money-service/routes"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
)

func initSpendingTypeForTest(app fiber.Router, db config.DummySpendingInstance) {

	spendingRepo := repositories.SpendingTypeRepositoryTest(db)
	spendingService := services.SpendingTypeService(spendingRepo)
	spendingController := controllers.SpendingTypeContoller(spendingService)
	routes.SpendingTypeController(app, spendingController)
}
func InitialTest(app fiber.Router, db config.DummySpendingInstance) {
	initSpendingTypeForTest(app, db)
}
