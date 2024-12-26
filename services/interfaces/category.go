package interfaces

import (
	"context"
	"sushi-backend/types/requests"
	"sushi-backend/types/responses"
)

type ICategoryService interface {
	GetAll(ctx context.Context, limit, offset int) *responses.Response
	GetById(ctx context.Context, id string) *responses.Response
	Create(ctx context.Context, request requests.CreateCategoryRequest) *responses.Response
	DeleteById(ctx context.Context, id string) *responses.Response
	UpdateById(ctx context.Context, id string, request requests.UpdateCategoryRequest) *responses.Response
}
