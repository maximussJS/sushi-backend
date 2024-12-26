package services

import (
	"context"
	"fmt"
	"sushi-backend/repositories/interfaces"
	"sushi-backend/services/dependecies"
	"sushi-backend/types/requests"
	"sushi-backend/types/responses"
)

type ProductService struct {
	productRepository  interfaces.IProductRepository
	categoryRepository interfaces.ICategoryRepository
}

func NewProductService(deps dependencies.ProductServiceDependencies) *ProductService {
	return &ProductService{
		productRepository:  deps.ProductRepository,
		categoryRepository: deps.CategoryRepository,
	}
}

func (p *ProductService) GetAll(ctx context.Context, limit, offset int) *responses.Response {
	products := p.productRepository.GetAll(ctx, limit, offset)

	return responses.NewSuccessResponse(products)
}

func (p *ProductService) GetById(ctx context.Context, id string) *responses.Response {
	product := p.productRepository.GetById(ctx, id)

	if product == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Product with id %s not found", id))
	}

	return responses.NewSuccessResponse(product)
}

func (p *ProductService) Create(ctx context.Context, request requests.CreateProductRequest) *responses.Response {
	if err := p.isValidCategoryId(ctx, request.CategoryId); err != nil {
		return err
	}

	existingProduct := p.productRepository.GetByName(ctx, request.Name)

	if existingProduct != nil {
		msg := fmt.Sprintf("Product with name %s already exists", request.Name)
		return responses.NewBadRequestResponse(msg)
	}

	productId := p.productRepository.Create(ctx, request.ToProductModel())

	newProduct := p.productRepository.GetById(ctx, productId)

	return responses.NewSuccessResponse(newProduct)
}

func (p *ProductService) UpdateById(ctx context.Context, id string, request requests.UpdateProductRequest) *responses.Response {
	product := p.productRepository.GetById(ctx, id)

	if product == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Product with id %s not found", id))
	}

	if request.CategoryId != "" {
		if err := p.isValidCategoryId(ctx, request.CategoryId); err != nil {
			return err
		}
	}

	p.productRepository.UpdateById(ctx, id, request.ToProductModel())

	updatedProduct := p.productRepository.GetById(ctx, id)

	return responses.NewSuccessResponse(updatedProduct)
}

func (p *ProductService) DeleteById(ctx context.Context, id string) *responses.Response {
	product := p.productRepository.GetById(ctx, id)

	if product == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Product with id %s not found", id))
	}

	p.productRepository.DeleteById(ctx, id)

	return responses.NewSuccessResponse(nil)
}

func (p *ProductService) isValidCategoryId(ctx context.Context, categoryId string) *responses.Response {
	category := p.categoryRepository.GetById(ctx, categoryId)

	if category == nil {
		return responses.NewBadRequestResponse(fmt.Sprintf("Category with id %s not found", categoryId))
	}

	return nil
}
