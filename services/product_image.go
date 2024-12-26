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

func (p *ProductImageService) GetById(ctx context.Context, id string) *responses.Response {
	product := p.imageRepository.GetById(ctx, id)

	if product == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Product with id %s not found", id))
	}

	return responses.NewSuccessResponse(product)
}

func (p *ProductImageService) Create(ctx context.Context, productId string, file multipart.File, header *multipart.FileHeader) *responses.Response {
	if err := p.isValidProductId(ctx, productId); err != nil {
		return err
	}

	publicId, secureUrl := p.cloudinary.Upload(ctx, file, header)

	imageId := p.imageRepository.Create(ctx, models.ProductImageModel{
		ProductId:          productId,
		CloudinaryPublicId: publicId,
		Url:                secureUrl,
	})

	newImage := p.imageRepository.GetById(ctx, imageId)

	return responses.NewSuccessResponse(newImage)
}

func (p *ProductImageService) DeleteById(ctx context.Context, id string) *responses.Response {
	image := p.imageRepository.GetById(ctx, id)

	if image == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Product image with id %s not found", id))
	}

	p.cloudinary.Delete(context.Background(), image.CloudinaryPublicId)

	p.imageRepository.DeleteById(ctx, id)

	return responses.NewSuccessResponse(nil)
}

func (p *ProductImageService) isValidProductId(ctx context.Context, productId string) *responses.Response {
	product := p.productRepository.GetById(ctx, productId)

	if product == nil {
		return responses.NewBadRequestResponse(fmt.Sprintf("Product with id %s not found", productId))
	}

	return nil
}
