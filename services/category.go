package services

import (
	"fmt"
	"sushi-backend/repositories/interfaces"
	"sushi-backend/services/dependecies"
	"sushi-backend/types/requests"
	"sushi-backend/types/responses"
	"sushi-backend/utils"
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
	categories := utils.PanicIfErrorWithResultReturning(c.categoryRepository.GetAll(limit, offset))

	return responses.NewSuccessResponse(categories)
}

func (c *CategoryService) GetById(id string) *responses.Response {
	category := utils.PanicIfErrorWithResultReturning(c.categoryRepository.GetById(id))

	if category == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Category with id %s not found", id))
	}

	return responses.NewSuccessResponse(category)
}

func (c *CategoryService) Create(request requests.CreateCategoryRequest) *responses.Response {
	existingCategory := utils.PanicIfErrorWithResultReturning(c.categoryRepository.GetByName(request.Name))

	if existingCategory != nil {
		msg := fmt.Sprintf("Category with name %s already exists", request.Name)
		return responses.NewBadRequestResponse(msg)
	}

	categoryId := utils.PanicIfErrorWithResultReturning(c.categoryRepository.Create(request.ToCategoryModel()))

	newCategory := utils.PanicIfErrorWithResultReturning(c.categoryRepository.GetById(categoryId))

	return responses.NewSuccessResponse(newCategory)
}

func (c *CategoryService) UpdateById(id string, request requests.UpdateCategoryRequest) *responses.Response {
	category := utils.PanicIfErrorWithResultReturning(c.categoryRepository.GetById(id))

	if category == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Category with id %s not found", id))
	}

	utils.PanicIfError(c.categoryRepository.UpdateById(id, request.ToCategoryModel()))

	updatedCategory := utils.PanicIfErrorWithResultReturning(c.categoryRepository.GetById(id))

	return responses.NewSuccessResponse(updatedCategory)
}

func (c *CategoryService) DeleteById(id string) *responses.Response {
	category := utils.PanicIfErrorWithResultReturning(c.categoryRepository.GetById(id))

	if category == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Category with id %s not found", id))
	}

	utils.PanicIfError(c.categoryRepository.DeleteById(id))

	return responses.NewSuccessResponse(nil)
}
