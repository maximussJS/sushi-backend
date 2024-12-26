package interfaces

import (
	"context"
	"mime/multipart"
	"sushi-backend/types/responses"
)

type IProductImageService interface {
	Create(ctx context.Context, productId string, file multipart.File, header *multipart.FileHeader) *responses.Response
	GetById(ctx context.Context, id string) *responses.Response
	DeleteById(ctx context.Context, id string) *responses.Response
}
