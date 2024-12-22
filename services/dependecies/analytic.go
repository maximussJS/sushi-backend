package dependencies

import (
	"go.uber.org/dig"
	"sushi-backend/config"
	"sushi-backend/repositories/interfaces"
)

type AnalyticServiceDependencies struct {
	dig.In

	Config          config.IConfig              `name:"Config"`
	OrderRepository interfaces.IOrderRepository `name:"OrderRepository"`
}
