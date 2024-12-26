package interfaces

import (
	"sushi-backend/models"
)

type ICategoryRepository interface {
	GetAll(limit, offset int) []models.CategoryModel
	GetByName(name string) *models.CategoryModel
	GetById(id string) *models.CategoryModel
	Create(category models.CategoryModel) string
	DeleteById(id string)
	UpdateById(id string, category models.CategoryModel)
}
