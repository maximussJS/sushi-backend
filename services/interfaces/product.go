package interfaces

import (
	"sushi-backend/types/requests"
	"sushi-backend/types/responses"
)

type IProductService interface {
	GetAll(limit, offset int) *responses.Response
	Create(request requests.CreateProductRequest) *responses.Response
	GetById(id string) *responses.Response
	DeleteById(id string) *responses.Response
	UpdateById(id string, request requests.UpdateProductRequest) *responses.Response
}
