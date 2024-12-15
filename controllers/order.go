package controllers

import (
	"net/http"
	"sushi-backend/controllers/dependencies"
	"sushi-backend/pkg/logger"
	"sushi-backend/types/responses"
)

type OrderController struct {
	logger logger.ILogger
}

func NewOrderController(deps dependencies.OrderHandlerDependencies) *OrderController {
	return &OrderController{
		logger: deps.Logger,
	}
}

func (h *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) *responses.Response {
	return responses.NewSuccessResponse(nil)
}
