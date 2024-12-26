package controllers

import (
	"net/http"
	"sushi-backend/controllers/dependencies"
	"sushi-backend/internal/logger"
	"sushi-backend/services/interfaces"
	"sushi-backend/types/responses"
	"sushi-backend/utils"
)

type AnalyticController struct {
	logger          logger.ILogger
	analyticService interfaces.IAnalyticService
}

func NewAnalyticController(deps dependencies.AnalyticControllerDependencies) *AnalyticController {
	return &AnalyticController{
		logger:          deps.Logger,
		analyticService: deps.AnalyticService,
	}
}

func (h *AnalyticController) GetOrdersAnalytic(_ http.ResponseWriter, r *http.Request) *responses.Response {
	startTimeInMs, err := utils.GetStartTimeInMsParam(r)

	if err != nil {
		return err
	}

	return h.analyticService.GetOrdersAnalytic(r.Context(), startTimeInMs)
}

func (h *AnalyticController) GetTopOrderedProducts(_ http.ResponseWriter, r *http.Request) *responses.Response {
	startTimeInMs, err := utils.GetStartTimeInMsParam(r)

	if err != nil {
		return err
	}

	limit, err := utils.GetLimitQueryParam(r, 100)

	if err != nil {
		return err
	}

	return h.analyticService.GetTopOrderedProducts(r.Context(), startTimeInMs, limit)
}
