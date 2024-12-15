package interfaces

import (
	"net/http"
	"sushi-backend/types/responses"
)

type IOrderController interface {
	CreateOrder(w http.ResponseWriter, r *http.Request) *responses.Response
}
