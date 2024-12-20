package db

import (
	"context"
	"go.uber.org/dig"
	"sushi-backend/config"
	"sushi-backend/internal/logger"
)

type DbDependencies struct {
	dig.In

	ShutdownContext context.Context `name:"ShutdownContext"`
	Logger          logger.ILogger  `name:"Logger"`
	Config          config.IConfig  `name:"Config"`
}
