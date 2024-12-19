package requests

import "sushi-backend/models"

type UpdateCategoryRequest struct {
	Name        *string `json:"name" validate:"omitempty,min=3,max=255"`
	Description *string `json:"description" validate:"omitempty,max=1024"`
}

func (r UpdateCategoryRequest) ToCategoryModel() models.CategoryModel {
	category := models.CategoryModel{}

	if r.Name != nil {
		category.Name = *r.Name
	}

	if r.Description != nil {
		category.Description = *r.Description
	}

	return category
}
