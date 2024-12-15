package services

import (
	"sushi-backend/repositories/interfaces"
	"sushi-backend/services/dependecies"
	"sushi-backend/types/responses"
)

type ProductService struct {
	productRepository interfaces.IProductRepository
}

func NewProductService(deps dependencies.ProductServiceDependencies) *ProductService {
	return &ProductService{
		productRepository: deps.ProductRepository,
	}
}

func (c *ProductService) GetAll() *responses.Response {
	products, err := c.productRepository.GetAll()

	if err != nil {
		return responses.NewInternalServerErrorResponse(err.Error())
	}

	return responses.NewSuccessResponse(products)
}
