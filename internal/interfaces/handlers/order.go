package handlers

import "net/http"

type IOrderHandler interface {
	CreateOrder(w http.ResponseWriter, r *http.Request)
}
