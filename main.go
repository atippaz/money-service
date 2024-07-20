package main

import (
	Config "money-service/config"
	"money-service/initials"
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
