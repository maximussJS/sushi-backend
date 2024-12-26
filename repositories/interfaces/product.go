package interfaces

import (
	"context"
	"sushi-backend/models"
)

type IProductRepository interface {
	GetAll(ctx context.Context, limit, offset int) []models.ProductModel
	GetByName(ctx context.Context, name string) *models.ProductModel
	GetById(ctx context.Context, id string) *models.ProductModel
	GetByIds(ctx context.Context, ids []string) []models.ProductModel
	Create(ctx context.Context, category models.ProductModel) string
	DeleteById(ctx context.Context, id string)
	UpdateById(ctx context.Context, id string, category models.ProductModel)
}
