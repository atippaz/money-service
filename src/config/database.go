package config

import (
	"fmt"
	"log"
	"money-service/src/entities"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"
)

func ConnectGormDb() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),

		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			// IgnoreRecordNotFoundError: true,
			// ParameterizedQueries:      true,
			Colorful: true,
		},
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ", cfg.HOST, cfg.USER, cfg.PASSWORD, cfg.DBNAME, cfg.DB_PORT)
	Database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&entities.SpendingTypeEntity{})
	Database.AutoMigrate(&entities.CustomTagEntity{})
	Database.AutoMigrate(&entities.ExpensesEntity{})
	Database.AutoMigrate(&entities.IncomesEntity{})
	Database.AutoMigrate(&entities.SystemTagEntity{})
	Database.AutoMigrate(&entities.UserEntity{})

	return Database
}

func ConnectTestDB() DummySpendingInstance {
	return GetInstanceDb()
}
