package requests

import (
	"sushi-backend/models"
)

type CreateCategoryRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=255"`
	Description string `json:"description" validate:"omitempty,max=1024"`
}

func (r CreateCategoryRequest) ToCategoryModel() models.CategoryModel {
	return models.CategoryModel{
		Name:        r.Name,
		Description: r.Description,
	}
}
