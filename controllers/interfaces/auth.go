package interfaces

import (
	"net/http"
	"sushi-backend/types/responses"
)

type IAuthController interface {
	Authorize(w http.ResponseWriter, r *http.Request) *responses.Response
	Verify(w http.ResponseWriter, r *http.Request) *responses.Response
	Refresh(w http.ResponseWriter, r *http.Request) *responses.Response
}
