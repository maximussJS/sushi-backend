package interfaces

import (
	"sushi-backend/models"
)

type IProductRepository interface {
	GetAll(limit, offset int) []models.ProductModel
	GetByName(name string) *models.ProductModel
	GetById(id string) *models.ProductModel
	GetByIds(ids []string) []models.ProductModel
	Create(category models.ProductModel) string
	DeleteById(id string)
	UpdateById(id string, category models.ProductModel)
}
