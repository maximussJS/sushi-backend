package services

import (
	"sushi-backend/repositories/interfaces"
	dependencies "sushi-backend/services/dependecies"
	"sushi-backend/types/responses"
	"sushi-backend/utils"
	"time"
)

type AnalyticService struct {
	orderRepository interfaces.IOrderRepository
}

func NewAnalyticService(deps dependencies.AnalyticServiceDependencies) *AnalyticService {
	return &AnalyticService{
		orderRepository: deps.OrderRepository,
	}
}

func (service *AnalyticService) GetOrdersAnalytic(startTimeInMs uint) *responses.Response {
	startTime, err := service.getStartTime(startTimeInMs)

	if err != nil {
		return err
	}

	orderAnalytic := utils.PanicIfErrorWithResultReturning(service.orderRepository.GetDeliveredOrdersAnalytic(startTime))

	orderAnalytic.StartTime = startTime.Format("2006-01-02 15:04:05")

	return responses.NewSuccessResponse(orderAnalytic)
}

func (service *AnalyticService) GetTopOrderedProducts(startTimeInMs uint, limit int) *responses.Response {
	startTime, err := service.getStartTime(startTimeInMs)

	if err != nil {
		return err
	}

	topOrderedProducts := utils.PanicIfErrorWithResultReturning(service.orderRepository.GetTopOrderedProducts(startTime, limit))

	for i := range topOrderedProducts {
		topOrderedProducts[i].StartTime = startTime.Format("2006-01-02 15:04:05")
	}

	return responses.NewSuccessResponse(topOrderedProducts)
}

func (service *AnalyticService) getStartTime(startTimeInMs uint) (time.Time, *responses.Response) {
	startTime := time.Unix(0, int64(startTimeInMs)*int64(time.Millisecond))

	if startTime.IsZero() {
		return time.Time{}, responses.NewBadRequestResponse("Invalid start time")
	}

	if startTime.After(time.Now()) {
		return time.Time{}, responses.NewBadRequestResponse("Start time cannot be in the future")
	}

	if startTime.Before(time.Now().AddDate(0, 0, -365)) {
		return time.Time{}, responses.NewBadRequestResponse("Start time cannot be older than 1 year")
	}

	return time.Unix(0, int64(startTimeInMs)*int64(time.Millisecond)), nil
}
