package controllers

import (
	"net/http"
	"sushi-backend/controllers/dependencies"
	"sushi-backend/internal/logger"
	"sushi-backend/services/interfaces"
	"sushi-backend/types/responses"
	"sushi-backend/utils"
)

type OrderFlowController struct {
	logger           logger.ILogger
	orderFlowService interfaces.IOrderFlowService
}

func NewOrderFlowController(deps dependencies.OrderFlowControllerDependencies) *OrderFlowController {
	return &OrderFlowController{
		logger:           deps.Logger,
		orderFlowService: deps.OrderFlowService,
	}
}

func (h *OrderFlowController) StartProcessing(_ http.ResponseWriter, r *http.Request) *responses.Response {
	id, err := utils.GetUIntIdParam(r)

	if err != nil {
		return err
	}

	estimatedTimeInMs, err := utils.GetEstimatedTimeInMsParam(r)

	if err != nil {
		return err
	}

	return h.orderFlowService.StartProcessing(id, estimatedTimeInMs)
}

func (h *OrderFlowController) ReadyToDeliver(_ http.ResponseWriter, r *http.Request) *responses.Response {
	id, err := utils.GetUIntIdParam(r)

	if err != nil {
		return err
	}

	return h.orderFlowService.ReadyToDeliver(id)
}

func (h *OrderFlowController) StartDelivering(_ http.ResponseWriter, r *http.Request) *responses.Response {
	id, err := utils.GetUIntIdParam(r)

	if err != nil {
		return err
	}

	estimatedTimeInMs, err := utils.GetEstimatedTimeInMsParam(r)

	if err != nil {
		return err
	}

	return h.orderFlowService.StartDelivering(id, estimatedTimeInMs)
}

func (h *OrderFlowController) Delivered(_ http.ResponseWriter, r *http.Request) *responses.Response {
	id, err := utils.GetUIntIdParam(r)

	if err != nil {
		return err
	}

	return h.orderFlowService.Delivered(id)
}

func (h *OrderFlowController) Cancel(_ http.ResponseWriter, r *http.Request) *responses.Response {
	id, err := utils.GetUIntIdParam(r)

	if err != nil {
		return err
	}

	return h.orderFlowService.Cancel(id)
}
