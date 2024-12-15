package interfaces

import (
	"sushi-backend/models"
)

type ICategoryRepository interface {
	GetAll(limit, offset int) ([]models.CategoryModel, error)
	Create(category models.CategoryModel) (string, error)
	FindByName(name string) (*models.CategoryModel, error)
	FindById(id string) (*models.CategoryModel, error)
}
