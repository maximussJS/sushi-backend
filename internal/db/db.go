package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB(deps DbDependecies) *gorm.DB {
	dsn := deps.Config.PostgresDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		deps.Logger.Error(fmt.Sprintf("Failed to connect to database: %s", err))
		panic(err)
	}

	deps.Logger.Log("Connected to database")

	return db
}
