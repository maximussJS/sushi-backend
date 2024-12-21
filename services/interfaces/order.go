package interfaces

import (
	"sushi-backend/types/requests"
	"sushi-backend/types/responses"
)

type IOrderService interface {
	Create(request requests.CreateOrderRequest) *responses.Response
	GetById(id uint) *responses.Response
	GetAll(limit, offset int) *responses.Response
	DeleteById(id uint) *responses.Response
}
