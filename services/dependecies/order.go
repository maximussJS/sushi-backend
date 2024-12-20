package dependencies

import (
	"go.uber.org/dig"
	"sushi-backend/internal/telegram"
)

type OrderServiceDependencies struct {
	dig.In

	Telegram telegram.ITelegram `name:"Telegram"`
}
