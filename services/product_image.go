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
	"sushi-backend/utils"
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
	product := utils.PanicIfErrorWithResultReturning(p.imageRepository.GetById(id))

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

	imageId := utils.PanicIfErrorWithResultReturning(p.imageRepository.Create(models.ProductImageModel{
		ProductId:          productId,
		CloudinaryPublicId: publicId,
		Url:                secureUrl,
	}))

	newImage := utils.PanicIfErrorWithResultReturning(p.imageRepository.GetById(imageId))

	return responses.NewSuccessResponse(newImage)
}

func (p *ProductImageService) DeleteById(id string) *responses.Response {
	image := utils.PanicIfErrorWithResultReturning(p.imageRepository.GetById(id))

	if image == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Product image with id %s not found", id))
	}

	p.cloudinary.Delete(context.Background(), image.CloudinaryPublicId)

	utils.PanicIfError(p.imageRepository.DeleteById(id))

	return responses.NewSuccessResponse(nil)
}

func (p *ProductImageService) isValidProductId(productId string) *responses.Response {
	product := utils.PanicIfErrorWithResultReturning(p.productRepository.GetById(productId))

	if product == nil {
		return responses.NewBadRequestResponse(fmt.Sprintf("Product with id %s not found", productId))
	}

	return nil
}
