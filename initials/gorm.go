package initials

import (
	"money-service/controllers"
	repositories "money-service/repositories/spending_type"
	"money-service/routes"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func initSpendingTypeForGorm(app fiber.Router, db_gorm *gorm.DB) {
	spendingRepo := repositories.SpendingTypeRepositoryGorm(db_gorm)
	spendingService := services.SpendingTypeService(spendingRepo)
	spendingController := controllers.SpendingTypeContoller(spendingService)
	routes.SpendingTypeController(app, spendingController)
}
func InitialGorm(app fiber.Router, db_gorm *gorm.DB) {
	initSpendingTypeForGorm(app, db_gorm)
}
