package dependencies

import (
	"go.uber.org/dig"
	"sushi-backend/internal/logger"
	"sushi-backend/services/interfaces"
)

type OrderFlowControllerDependencies struct {
	dig.In

	Logger           logger.ILogger               `name:"Logger"`
	OrderFlowService interfaces.IOrderFlowService `name:"OrderFlowService"`
}
