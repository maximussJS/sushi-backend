package dependencies

import (
	"go.uber.org/dig"
	"sushi-backend/internal/logger"
	"sushi-backend/services/interfaces"
)

type AnalyticControllerDependencies struct {
	dig.In

	Logger          logger.ILogger              `name:"Logger"`
	AnalyticService interfaces.IAnalyticService `name:"AnalyticService"`
}
