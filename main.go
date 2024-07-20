package main

import (
	Config "money-service/src/config"
	"money-service/src/initials"
)

func main() {
	app, config := Config.GetInstanceServer()

	// db_test := Config.ConnectTestDB()
	// initials.InitialTest(db_test)

	db_gorm := Config.ConnectGormDb()
	initials.InitialGorm(db_gorm)

	initials.InitApplication(app, config)

	Config.StartServer()
}
