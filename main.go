package main

import (
	Config "money-service/config"
	"money-service/initials"
)

func main() {
	db_gorm := Config.ConnectGormDb()
	// db_test := Config.ConnectTestDB()
	app := Config.GetInstanceServer()
	// initSpendingTypeForTest(app, db_test)
	initials.InitialGorm(app, db_gorm)

	Config.StartServer()
}
