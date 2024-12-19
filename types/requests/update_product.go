package requests

import "sushi-backend/models"

type UpdateProductRequest struct {
	Name        string  `json:"name" validate:"omitempty,min=1,max=255"`
	Description string  `json:"description,omitempty" validate:"omitempty,max=1024"`
	Price       float64 `json:"price" validate:"omitempty,gt=0"`
	CategoryId  string  `json:"categoryId" validate:"omitempty"`
}

func (r UpdateProductRequest) ToProductModel() models.ProductModel {
	product := models.ProductModel{}

	if r.Name != "" {
		product.Name = r.Name
	}

	if r.Description != "" {
		product.Description = r.Description
	}

	if r.Price != 0 {
		product.Price = r.Price
	}

	if r.CategoryId != "" {
		product.CategoryId = r.CategoryId
	}

	return product
}
