package services

import (
	"context"
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

func (c *CategoryService) GetAll(ctx context.Context, limit, offset int) *responses.Response {
	categories := c.categoryRepository.GetAll(ctx, limit, offset)

	return responses.NewSuccessResponse(categories)
}

func (c *CategoryService) GetById(ctx context.Context, id string) *responses.Response {
	category := c.categoryRepository.GetById(ctx, id)

	if category == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Category with id %s not found", id))
	}

	return responses.NewSuccessResponse(category)
}

func (c *CategoryService) Create(ctx context.Context, request requests.CreateCategoryRequest) *responses.Response {
	existingCategory := c.categoryRepository.GetByName(ctx, request.Name)

	if existingCategory != nil {
		msg := fmt.Sprintf("Category with name %s already exists", request.Name)
		return responses.NewBadRequestResponse(msg)
	}

	categoryId := c.categoryRepository.Create(ctx, request.ToCategoryModel())

	newCategory := c.categoryRepository.GetById(ctx, categoryId)

	return responses.NewSuccessResponse(newCategory)
}

func (c *CategoryService) UpdateById(ctx context.Context, id string, request requests.UpdateCategoryRequest) *responses.Response {
	category := c.categoryRepository.GetById(ctx, id)

	if category == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Category with id %s not found", id))
	}

	c.categoryRepository.UpdateById(ctx, id, request.ToCategoryModel())

	updatedCategory := c.categoryRepository.GetById(ctx, id)

	return responses.NewSuccessResponse(updatedCategory)
}

func (c *CategoryService) DeleteById(ctx context.Context, id string) *responses.Response {
	category := c.categoryRepository.GetById(ctx, id)

	if category == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Category with id %s not found", id))
	}

	c.categoryRepository.DeleteById(ctx, id)

	return responses.NewSuccessResponse(nil)
}
