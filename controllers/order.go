package controllers

import (
	"net/http"
	"sushi-backend/controllers/dependencies"
	"sushi-backend/internal/logger"
	"sushi-backend/services/interfaces"
	"sushi-backend/types/responses"
)

type OrderController struct {
	logger       logger.ILogger
	orderService interfaces.IOrderService
}

func NewOrderController(deps dependencies.OrderHandlerDependencies) *OrderController {
	return &OrderController{
		logger:       deps.Logger,
		orderService: deps.OrderService,
	}
}

func (h *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) *responses.Response {
	return h.orderService.GetById("test")
}
