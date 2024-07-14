package initials

import (
	"money-service/controllers"
	spending_type "money-service/repositories/spending_type"
	user "money-service/repositories/user"

	"money-service/routes"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func initSpendingTypeForGorm(app fiber.Router, db_gorm *gorm.DB) {
	spendingRepo := spending_type.SpendingTypeRepositoryGorm(db_gorm)
	spendingService := services.SpendingTypeService(spendingRepo)
	spendingController := controllers.SpendingTypeContoller(spendingService)
	routes.SpendingTypeRoute(app, spendingController)
}
func initUserForGorm(app fiber.Router, db_gorm *gorm.DB) {
	spendingRepo := user.UserRepositoryGorm(db_gorm)
	spendingService := services.UserService(spendingRepo)
	spendingController := controllers.UserContoller(spendingService)
	routes.UserRoute(app, spendingController)
}
func InitialGorm(app fiber.Router, db_gorm *gorm.DB) {
	initSpendingTypeForGorm(app, db_gorm)
	initUserForGorm(app, db_gorm)
}
