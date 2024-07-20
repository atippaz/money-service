package main

import (
	Config "money-service/config"
	"money-service/initials"
)

// var UserId = uuid.MustParse("")

func main() {
	db_gorm := Config.ConnectGormDb()
	// db_test := Config.ConnectTestDB()
	app, config := Config.GetInstanceServer()
	initials.InitialGorm(db_gorm)
	// initSpendingTypeForTest(app, db_test)
	initials.InitApplication(app, config)

	Config.StartServer()
}
