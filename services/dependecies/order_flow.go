package dependencies

import (
	"go.uber.org/dig"
	"sushi-backend/config"
	"sushi-backend/internal/telegram"
	"sushi-backend/repositories/interfaces"
)

type OrderFlowServiceDependencies struct {
	dig.In

	Config          config.IConfig              `name:"Config"`
	Telegram        telegram.ITelegram          `name:"Telegram"`
	OrderRepository interfaces.IOrderRepository `name:"OrderRepository"`
}
