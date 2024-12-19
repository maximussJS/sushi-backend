package interfaces

import (
	"sushi-backend/types/requests"
	"sushi-backend/types/responses"
)

type ICategoryService interface {
	GetAll(limit, offset int) *responses.Response
	GetById(id string) *responses.Response
	Create(request requests.CreateCategoryRequest) *responses.Response
	DeleteById(id string) *responses.Response
	UpdateById(id string, request requests.UpdateCategoryRequest) *responses.Response
}
