package main

import (
	Config "money-service/config"
	"money-service/controllers"
	repositories "money-service/repositories/spending_type"
	"money-service/routes"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func main() {
	// db_gorm := databases.ConnectGormDb()
	db_test := Config.ConnectTestDB()
	app := Config.GetInstanceServer()
	initSpendingTypeForTest(app, db_test)
	Config.StartServer()
}
func initSpendingTypeForTest(app fiber.Router, db Config.DummySpendingInstance) {

	spendingRepo := repositories.SpendingTypeRepositoryTest(db)
	spendingService := services.SpendingTypeService(spendingRepo)
	spendingController := controllers.SpendingTypeContoller(spendingService)
	routes.SpendingTypeController(app, spendingController)
}
func initSpendingTypeForGorm(app fiber.Router, db_gorm *gorm.DB) {
	spendingRepo := repositories.SpendingTypeRepositoryGorm(db_gorm)
	spendingService := services.SpendingTypeService(spendingRepo)
	spendingController := controllers.SpendingTypeContoller(spendingService)
	routes.SpendingTypeController(app, spendingController)
}
