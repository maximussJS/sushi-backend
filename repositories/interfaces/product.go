package interfaces

import (
	"sushi-backend/models"
)

type IProductRepository interface {
	GetAll(limit, offset int) ([]models.ProductModel, error)
	GetByName(name string) (*models.ProductModel, error)
	GetById(id string) (*models.ProductModel, error)
	GetByIds(ids []string) ([]models.ProductModel, error)
	Create(category models.ProductModel) (string, error)
	DeleteById(id string) error
	UpdateById(id string, category models.ProductModel) error
}
