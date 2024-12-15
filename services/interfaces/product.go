package interfaces

import (
	"sushi-backend/types/responses"
)

type IProductService interface {
	GetAll() *responses.Response
}
