package main

import (
	Config "money-service/config"
	"money-service/initials"
	"money-service/middlewares"
)

// var UserId = uuid.MustParse("")

func main() {
	db_gorm := Config.ConnectGormDb()
	// db_test := Config.ConnectTestDB()
	app := Config.GetInstanceServer()
	initials.InitialGorm(db_gorm)
	// initSpendingTypeForTest(app, db_test)
	app.Use(middlewares.ApiKeyMiddleware())
	initials.InitApplication(app)

	Config.StartServer()
}
