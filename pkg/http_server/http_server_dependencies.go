package http_server

import (
	"context"
	"go.uber.org/dig"
	"sushi-backend/internal/interfaces/handlers"
	"sushi-backend/pkg/config"
	"sushi-backend/pkg/logger"
	"sushi-backend/pkg/rate_limit"
	"sync"
)

type HttpServerDependencies struct {
	dig.In

	ShutdownWaitGroup *sync.WaitGroup           `name:"ShutdownWaitGroup"`
	ShutdownContext   context.Context           `name:"ShutdownContext"`
	Logger            logger.ILogger            `name:"Logger"`
	Config            config.IConfig            `name:"Config"`
	OrderHandler      handlers.IOrderHandler    `name:"OrderHandler"`
	IPRateLimiter     rate_limit.IIpRateLimiter `name:"IpRateLimiter"`
}
