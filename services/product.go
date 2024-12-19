package services

import (
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

func (p *ProductService) GetAll(limit, offset int) *responses.Response {
	products, err := p.productRepository.GetAll(limit, offset)
	if err != nil {
		return responses.NewInternalServerErrorResponse(err.Error())
	}

	return responses.NewSuccessResponse(products)
}

func (p *ProductService) GetById(id string) *responses.Response {
	product, err := p.productRepository.FindById(id)
	if err != nil {
		return responses.NewInternalServerErrorResponse(err.Error())
	}

	if product == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Product with id %s not found", id))
	}

	return responses.NewSuccessResponse(product)
}

func (p *ProductService) Create(request requests.CreateProductRequest) *responses.Response {
	if err := p.isValidCategoryId(request.CategoryId); err != nil {
		return err
	}

	existingProduct, err := p.productRepository.FindByName(request.Name)
	if err != nil {
		return responses.NewInternalServerErrorResponse(err.Error())
	}

	if existingProduct != nil {
		msg := fmt.Sprintf("Product with name %s already exists", request.Name)
		return responses.NewBadRequestResponse(msg)
	}

	productId, err := p.productRepository.Create(request.ToProductModel())
	if err != nil {
		return responses.NewInternalServerErrorResponse(err.Error())
	}

	newProduct, err := p.productRepository.FindById(productId)
	if err != nil {
		return responses.NewInternalServerErrorResponse(err.Error())
	}

	return responses.NewSuccessResponse(newProduct)
}

func (p *ProductService) UpdateById(id string, request requests.UpdateProductRequest) *responses.Response {
	product, err := p.productRepository.FindById(id)
	if err != nil {
		return responses.NewInternalServerErrorResponse(err.Error())
	}

	if product == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Product with id %s not found", id))
	}

	if request.CategoryId != "" {
		if err := p.isValidCategoryId(request.CategoryId); err != nil {
			return err
		}
	}

	err = p.productRepository.UpdateById(id, request.ToProductModel())
	if err != nil {
		return responses.NewInternalServerErrorResponse(err.Error())
	}

	updatedProduct, err := p.productRepository.FindById(id)
	if err != nil {
		return responses.NewInternalServerErrorResponse(err.Error())
	}

	return responses.NewSuccessResponse(updatedProduct)
}

func (p *ProductService) DeleteById(id string) *responses.Response {
	product, err := p.productRepository.FindById(id)
	if err != nil {
		return responses.NewInternalServerErrorResponse(err.Error())
	}

	if product == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Product with id %s not found", id))
	}

	err = p.productRepository.DeleteById(id)
	if err != nil {
		return responses.NewInternalServerErrorResponse(err.Error())
	}

	return responses.NewSuccessResponse(nil)
}

func (p *ProductService) isValidCategoryId(categoryId string) *responses.Response {
	category, err := p.categoryRepository.FindById(categoryId)

	if err != nil {
		return responses.NewInternalServerErrorResponse(err.Error())
	}

	if category == nil {
		return responses.NewBadRequestResponse(fmt.Sprintf("Category with id %s not found", categoryId))
	}

	return nil
}
