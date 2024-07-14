package main

import (
	Config "money-service/config"
	"money-service/controllers"
	"money-service/gorm_repository"
	"money-service/other_db"
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
func initSpendingTypeForTest(app fiber.Router, db other_db.DummySpendingInstance) {

	spendingRepo := other_db.SpendingTypeRepositoryTest(db)
	spendingService := services.SpendingTypeService(spendingRepo)
	spendingController := controllers.SpendingTypeContoller(spendingService)
	routes.SpendingTypeController(app, spendingController)
}
func initSpendingTypeForGorm(app fiber.Router, db_gorm *gorm.DB) {
	spendingRepo := gorm_repository.SpendingTypeRepositoryGorm(db_gorm)
	spendingService := services.SpendingTypeService(spendingRepo)
	spendingController := controllers.SpendingTypeContoller(spendingService)
	routes.SpendingTypeController(app, spendingController)
}
