package dependencies

import (
	"go.uber.org/dig"
	"sushi-backend/pkg/logger"
)

type OrderHandlerDependencies struct {
	dig.In

	Logger logger.ILogger `name:"Logger"`
}
