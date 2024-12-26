package interfaces

import (
	"context"
	"sushi-backend/models"
)

type IProductImageRepository interface {
	Create(ctx context.Context, category models.ProductImageModel) string
	GetById(ctx context.Context, id string) *models.ProductImageModel
	DeleteById(ctx context.Context, id string)
}
