package interfaces

import (
	"sushi-backend/models"
)

type ICategoryRepository interface {
	GetAll(limit, offset int) ([]models.CategoryModel, error)
	GetByName(name string) (*models.CategoryModel, error)
	GetById(id string) (*models.CategoryModel, error)
	Create(category models.CategoryModel) (string, error)
	DeleteById(id string) error
	UpdateById(id string, category models.CategoryModel) error
}
