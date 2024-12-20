package services

import (
	"sushi-backend/internal/telegram"
	dependencies "sushi-backend/services/dependecies"
	"sushi-backend/types/responses"
)

type OrderService struct {
	telegram telegram.ITelegram
}

func NewOrderService(deps dependencies.OrderServiceDependencies) *OrderService {
	return &OrderService{
		telegram: deps.Telegram,
	}
}

func (o *OrderService) GetById(id string) *responses.Response {
	o.telegram.SendMessageToOrdersChannel("test")
	return responses.NewSuccessResponse(nil)
}
