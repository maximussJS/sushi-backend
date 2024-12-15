package interfaces

import (
	"sushi-backend/models"
)

type IProductRepository interface {
	GetAll() ([]models.ProductModel, error)
}
