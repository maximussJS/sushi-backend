package interfaces

import (
	"net/http"
	"sushi-backend/types/responses"
)

type ICategoryController interface {
	GetAll(w http.ResponseWriter, r *http.Request) *responses.Response
	Create(w http.ResponseWriter, r *http.Request) *responses.Response
	GetById(w http.ResponseWriter, r *http.Request) *responses.Response
	DeleteById(w http.ResponseWriter, r *http.Request) *responses.Response
	UpdateById(w http.ResponseWriter, r *http.Request) *responses.Response
}
