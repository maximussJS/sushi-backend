package interfaces

import (
	"mime/multipart"
	"sushi-backend/types/responses"
)

type IProductImageService interface {
	Create(productId string, file multipart.File, header *multipart.FileHeader) *responses.Response
	GetById(id string) *responses.Response
	DeleteById(id string) *responses.Response
}
