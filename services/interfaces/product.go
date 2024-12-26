package interfaces

import (
	"context"
	"sushi-backend/types/requests"
	"sushi-backend/types/responses"
)

type IProductService interface {
	GetAll(ctx context.Context, limit, offset int) *responses.Response
	Create(ctx context.Context, request requests.CreateProductRequest) *responses.Response
	GetById(ctx context.Context, id string) *responses.Response
	DeleteById(ctx context.Context, id string) *responses.Response
	UpdateById(ctx context.Context, id string, request requests.UpdateProductRequest) *responses.Response
}
