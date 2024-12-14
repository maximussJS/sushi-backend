package rate_limit

import (
	"context"
	"go.uber.org/dig"
	"sushi-backend/pkg/config"
	"sushi-backend/pkg/logger"
)

type IpRateLimiterDependencies struct {
	dig.In

	ShutdownContext context.Context `name:"ShutdownContext"`
	Logger          logger.ILogger  `name:"Logger"`
	Config          config.IConfig  `name:"Config"`
}
