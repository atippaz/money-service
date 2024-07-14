package main

import (
	Config "money-service/config"
	"money-service/initials"
)

func main() {
	db_gorm := Config.ConnectGormDb()
	// db_test := Config.ConnectTestDB()
	app := Config.GetInstanceServer()
	initials.InitialGorm(db_gorm)
	// initSpendingTypeForTest(app, db_test)
	initials.InitApplication(app)

	Config.StartServer()
}
