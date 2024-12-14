package handlers

import (
	"net/http"
	"sushi-backend/internal/dependencies/handlers"
	"sushi-backend/pkg/logger"
)

type OrderHandler struct {
	logger logger.ILogger
}

func NewOrderHandler(deps handlers.OrderHandlerDependencies) *OrderHandler {
	return &OrderHandler{
		logger: deps.Logger,
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("good"))
}
