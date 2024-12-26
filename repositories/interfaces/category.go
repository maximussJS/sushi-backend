package interfaces

import (
	"context"
	"sushi-backend/models"
)

type ICategoryRepository interface {
	GetAll(ctx context.Context, limit, offset int) []models.CategoryModel
	GetByName(ctx context.Context, name string) *models.CategoryModel
	GetById(ctx context.Context, id string) *models.CategoryModel
	Create(ctx context.Context, category models.CategoryModel) string
	DeleteById(ctx context.Context, id string)
	UpdateById(ctx context.Context, id string, category models.CategoryModel)
}
