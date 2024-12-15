package interfaces

import (
	"net/http"
	"sushi-backend/types/responses"
)

type IProductController interface {
	GetAll(w http.ResponseWriter, r *http.Request) *responses.Response
}
