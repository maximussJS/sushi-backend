package controllers

import (
	"net/http"
	"sushi-backend/controllers/dependencies"
	"sushi-backend/internal/logger"
	"sushi-backend/services/interfaces"
	"sushi-backend/types/requests"
	"sushi-backend/types/responses"
	"sushi-backend/utils"
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

func (h *OrderController) Create(w http.ResponseWriter, r *http.Request) *responses.Response {
	var req requests.CreateOrderRequest

	err := utils.DecodeJSONBody(w, r, &req)
	if err != nil {
		return err
	}

	return h.orderService.Create(req)
}

func (h *OrderController) GetById(_ http.ResponseWriter, r *http.Request) *responses.Response {
	id, err := utils.GetUIntIdParam(r)

	if err != nil {
		return err
	}

	return h.orderService.GetById(id)
}

func (h *OrderController) GetAll(_ http.ResponseWriter, r *http.Request) *responses.Response {
	limit, err := utils.GetLimitQueryParam(r, 100)

	if err != nil {
		return err
	}

	offset, err := utils.GetOffsetQueryParam(r, 0)

	if err != nil {
		return err
	}

	return h.orderService.GetAll(limit, offset)
}

func (h *OrderController) DeleteById(_ http.ResponseWriter, r *http.Request) *responses.Response {
	id, err := utils.GetUIntIdParam(r)

	if err != nil {
		return err
	}

	return h.orderService.DeleteById(id)
}
