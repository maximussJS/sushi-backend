package interfaces

import (
	"net/http"
	"sushi-backend/types/responses"
)

type IOrderFlowController interface {
	StartProcessing(w http.ResponseWriter, r *http.Request) *responses.Response
	ReadyToDeliver(w http.ResponseWriter, r *http.Request) *responses.Response
	StartDelivering(w http.ResponseWriter, r *http.Request) *responses.Response
	Delivered(w http.ResponseWriter, r *http.Request) *responses.Response
	Cancel(w http.ResponseWriter, r *http.Request) *responses.Response
}
