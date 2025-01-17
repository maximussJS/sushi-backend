package dependencies

import (
	"go.uber.org/dig"
	"sushi-backend/internal/logger"
	"sushi-backend/services/interfaces"
)

type OrderControllerDependencies struct {
	dig.In

	Logger       logger.ILogger           `name:"Logger"`
	OrderService interfaces.IOrderService `name:"OrderService"`
}
