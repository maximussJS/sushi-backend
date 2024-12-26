package interfaces

import "sushi-backend/models"

type IProductImageRepository interface {
	Create(category models.ProductImageModel) string
	GetById(id string) *models.ProductImageModel
	DeleteById(id string)
}
