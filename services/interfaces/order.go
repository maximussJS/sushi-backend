package interfaces

import (
	"context"
	"sushi-backend/types/requests"
	"sushi-backend/types/responses"
)

type IOrderService interface {
	Create(ctx context.Context, request requests.CreateOrderRequest) *responses.Response
	GetById(ctx context.Context, id uint) *responses.Response
	GetAll(ctx context.Context, limit, offset int) *responses.Response
	DeleteById(ctx context.Context, id uint) *responses.Response
}
