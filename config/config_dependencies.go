package config

import (
	"go.uber.org/dig"
	"sushi-backend/internal/logger"
)

type ConfigDependencies struct {
	dig.In

	Logger logger.ILogger `name:"Logger"`
}
