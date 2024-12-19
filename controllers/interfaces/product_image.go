package interfaces

import (
	"net/http"
	"sushi-backend/types/responses"
)

type IProductImageController interface {
	Create(w http.ResponseWriter, r *http.Request) *responses.Response
	GetById(w http.ResponseWriter, r *http.Request) *responses.Response
	DeleteById(w http.ResponseWriter, r *http.Request) *responses.Response
}