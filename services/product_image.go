package services

import (
	"context"
	"fmt"
	"mime/multipart"
	"sushi-backend/internal/cloudinary"
	"sushi-backend/models"
	"sushi-backend/repositories/interfaces"
	"sushi-backend/services/dependecies"
	"sushi-backend/types/responses"
)

type ProductImageService struct {
	cloudinary        cloudinary.ICloudinary
	productRepository interfaces.IProductRepository
	imageRepository   interfaces.IProductImageRepository
}

func NewProductImageService(deps dependencies.ProductImageServiceDependencies) *ProductImageService {
	return &ProductImageService{
		cloudinary:        deps.Cloudinary,
		productRepository: deps.ProductRepository,
		imageRepository:   deps.ProductImageRepository,
	}
}

func (p *ProductImageService) GetById(id string) *responses.Response {
	product, err := p.imageRepository.FindById(id)
	if err != nil {
		return responses.NewInternalServerErrorResponse(err.Error())
	}

	if product == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Product with id %s not found", id))
	}

	return responses.NewSuccessResponse(product)
}

func (p *ProductImageService) Create(productId string, file multipart.File, header *multipart.FileHeader) *responses.Response {
	if err := p.isValidProductId(productId); err != nil {
		return err
	}

	ctx := context.Background()
	publicId, secureUrl := p.cloudinary.Upload(ctx, file, header)

	imageId, err := p.imageRepository.Create(models.ProductImageModel{
		ProductId:          productId,
		CloudinaryPublicId: publicId,
		Url:                secureUrl,
	})

	if err != nil {
		return responses.NewInternalServerErrorResponse(err.Error())
	}

	newImage, err := p.imageRepository.FindById(imageId)
	if err != nil {
		return responses.NewInternalServerErrorResponse(err.Error())
	}

	return responses.NewSuccessResponse(newImage)
}

func (p *ProductImageService) DeleteById(id string) *responses.Response {
	image, err := p.imageRepository.FindById(id)
	if err != nil {
		return responses.NewInternalServerErrorResponse(err.Error())
	}

	if image == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Product image with id %s not found", id))
	}

	p.cloudinary.Delete(context.Background(), image.CloudinaryPublicId)

	err = p.imageRepository.DeleteById(id)
	if err != nil {
		return responses.NewInternalServerErrorResponse(err.Error())
	}

	return responses.NewSuccessResponse(nil)
}

func (p *ProductImageService) isValidProductId(productId string) *responses.Response {
	product, err := p.productRepository.FindById(productId)

	if err != nil {
		return responses.NewInternalServerErrorResponse(err.Error())
	}

	if product == nil {
		return responses.NewBadRequestResponse(fmt.Sprintf("Product with id %s not found", productId))
	}

	return nil
}
