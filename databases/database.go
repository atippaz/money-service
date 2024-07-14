package databases

import (
	"fmt"
	"money-service/config"
	"money-service/entities"
	"money-service/other_db"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func ConnectGormDb() *gorm.DB {
	cfg := config.LoadConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ", cfg.HOST, cfg.USER, cfg.PASSWORD, cfg.DBNAME, cfg.DB_PORT)
	fmt.Print(dsn)
	Database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Database.AutoMigrate(&models.User{})
	Database.AutoMigrate(&entities.SpendingTypeEntity{})
	return Database
}

func ConnectTestDB() other_db.DummySpendingInstance {
	return other_db.GetInstanceDb()
}
