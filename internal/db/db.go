package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sushi-backend/constants"
)

func NewDB(deps DbDependencies) *gorm.DB {
	dsn := deps.Config.PostgresDSN()

	logMode := logger.Info

	if deps.Config.AppEnv() == constants.ProductionEnv {
		logMode = logger.Error
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})

	if err != nil {
		deps.Logger.Error(fmt.Sprintf("Failed to connect to database: %s", err))
		panic(err)
	}

	deps.Logger.Log("Connected to database")

	return db
}
