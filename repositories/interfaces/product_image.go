package interfaces

import "sushi-backend/models"

type IProductImageRepository interface {
	Create(category models.ProductImageModel) (string, error)
	GetById(id string) (*models.ProductImageModel, error)
	DeleteById(id string) error
}
