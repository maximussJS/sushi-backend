package http_server

import (
	"context"
	"go.uber.org/dig"
	"sushi-backend/config"
	"sushi-backend/internal/logger"
	"sushi-backend/router"
	"sync"
)

type HttpServerDependencies struct {
	dig.In

	ShutdownWaitGroup *sync.WaitGroup `name:"ShutdownWaitGroup"`
	ShutdownContext   context.Context `name:"ShutdownContext"`
	Router            *router.Router  `name:"Router"`
	Logger            logger.ILogger  `name:"Logger"`
	Config            config.IConfig  `name:"Config"`
}
