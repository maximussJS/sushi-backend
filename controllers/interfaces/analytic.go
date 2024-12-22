package interfaces

import (
	"net/http"
	"sushi-backend/types/responses"
)

type IAnalyticController interface {
	GetOrdersAnalytic(w http.ResponseWriter, r *http.Request) *responses.Response
	GetTopOrderedProducts(w http.ResponseWriter, r *http.Request) *responses.Response
}
