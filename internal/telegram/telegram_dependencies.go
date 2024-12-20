package telegram

import (
	"go.uber.org/dig"
	"sushi-backend/config"
	"sushi-backend/internal/logger"
)

type TelegramDependencies struct {
	dig.In

	Logger logger.ILogger `name:"Logger"`
	Config config.IConfig `name:"Config"`
}
