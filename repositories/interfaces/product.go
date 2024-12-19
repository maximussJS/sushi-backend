package interfaces

import (
	"sushi-backend/models"
)

type IProductRepository interface {
	GetAll(limit, offset int) ([]models.ProductModel, error)
	Create(category models.ProductModel) (string, error)
	FindByName(name string) (*models.ProductModel, error)
	FindById(id string) (*models.ProductModel, error)
	DeleteById(id string) error
	UpdateById(id string, category models.ProductModel) error
}
