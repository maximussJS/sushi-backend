package interfaces

import "sushi-backend/types/responses"

type IOrderService interface {
	GetById(id string) *responses.Response
}
