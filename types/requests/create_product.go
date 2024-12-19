package requests

import "sushi-backend/models"

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required,min=1,max=255"`
	Description string  `json:"description,omitempty" validate:"omitempty,max=1024"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	CategoryId  string  `json:"categoryId" validate:"required"`
}

func (r CreateProductRequest) ToProductModel() models.ProductModel {
	return models.ProductModel{
		Name:        r.Name,
		Description: r.Description,
		Price:       r.Price,
		CategoryId:  r.CategoryId,
	}
}
