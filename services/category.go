package services

import (
	"fmt"
	"sushi-backend/repositories/interfaces"
	"sushi-backend/services/dependecies"
	"sushi-backend/types/requests"
	"sushi-backend/types/responses"
)

type CategoryService struct {
	categoryRepository interfaces.ICategoryRepository
}

func NewCategoryService(deps dependencies.CategoryServiceDependencies) *CategoryService {
	return &CategoryService{
		categoryRepository: deps.CategoryRepository,
	}
}

func (c *CategoryService) GetAll(limit, offset int) *responses.Response {
	categories, err := c.categoryRepository.GetAll(limit, offset)

	if err != nil {
		return responses.NewInternalServerErrorResponse(err.Error())
	}

	return responses.NewSuccessResponse(categories)
}

func (c *CategoryService) GetById(id string) *responses.Response {
	category, err := c.categoryRepository.FindById(id)

	if err != nil {
		return responses.NewInternalServerErrorResponse(err.Error())
	}

	if category == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Category with id %s not found", id))
	}

	return responses.NewSuccessResponse(category)
}

func (c *CategoryService) Create(request requests.CreateCategoryRequest) *responses.Response {
	existingCategory, err := c.categoryRepository.FindByName(request.Name)

	if err != nil {
		return responses.NewInternalServerErrorResponse(err.Error())
	}

	if existingCategory != nil {
		msg := fmt.Sprintf("Category with name %s already exists", request.Name)
		return responses.NewBadRequestResponse(msg)
	}

	categoryId, err := c.categoryRepository.Create(request.ToCategoryModel())

	if err != nil {
		return responses.NewInternalServerErrorResponse(err.Error())
	}

	newCategory, err := c.categoryRepository.FindById(categoryId)

	if err != nil {
		return responses.NewInternalServerErrorResponse(err.Error())
	}

	return responses.NewSuccessResponse(newCategory)
}
