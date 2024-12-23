package services

import (
	"fmt"
	"sushi-backend/repositories/interfaces"
	"sushi-backend/services/dependecies"
	"sushi-backend/types/requests"
	"sushi-backend/types/responses"
	"sushi-backend/utils"
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
	products := utils.PanicIfErrorWithResultReturning(p.productRepository.GetAll(limit, offset))

	return responses.NewSuccessResponse(products)
}

func (p *ProductService) GetById(id string) *responses.Response {
	product := utils.PanicIfErrorWithResultReturning(p.productRepository.GetById(id))

	if product == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Product with id %s not found", id))
	}

	return responses.NewSuccessResponse(product)
}

func (p *ProductService) Create(request requests.CreateProductRequest) *responses.Response {
	if err := p.isValidCategoryId(request.CategoryId); err != nil {
		return err
	}

	existingProduct := utils.PanicIfErrorWithResultReturning(p.productRepository.GetByName(request.Name))

	if existingProduct != nil {
		msg := fmt.Sprintf("Product with name %s already exists", request.Name)
		return responses.NewBadRequestResponse(msg)
	}

	productId := utils.PanicIfErrorWithResultReturning(p.productRepository.Create(request.ToProductModel()))

	newProduct := utils.PanicIfErrorWithResultReturning(p.productRepository.GetById(productId))

	return responses.NewSuccessResponse(newProduct)
}

func (p *ProductService) UpdateById(id string, request requests.UpdateProductRequest) *responses.Response {
	product := utils.PanicIfErrorWithResultReturning(p.productRepository.GetById(id))

	if product == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Product with id %s not found", id))
	}

	if request.CategoryId != "" {
		if err := p.isValidCategoryId(request.CategoryId); err != nil {
			return err
		}
	}

	utils.PanicIfError(p.productRepository.UpdateById(id, request.ToProductModel()))

	updatedProduct := utils.PanicIfErrorWithResultReturning(p.productRepository.GetById(id))

	return responses.NewSuccessResponse(updatedProduct)
}

func (p *ProductService) DeleteById(id string) *responses.Response {
	product := utils.PanicIfErrorWithResultReturning(p.productRepository.GetById(id))

	if product == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Product with id %s not found", id))
	}

	utils.PanicIfError(p.productRepository.DeleteById(id))

	return responses.NewSuccessResponse(nil)
}

func (p *ProductService) isValidCategoryId(categoryId string) *responses.Response {
	category := utils.PanicIfErrorWithResultReturning(p.categoryRepository.GetById(categoryId))

	if category == nil {
		return responses.NewBadRequestResponse(fmt.Sprintf("Category with id %s not found", categoryId))
	}

	return nil
}
