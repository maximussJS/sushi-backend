package interfaces

import (
	"sushi-backend/types/responses"
)

type IAnalyticService interface {
	GetOrdersAnalytic(startTimeInMs uint) *responses.Response
	GetTopOrderedProducts(startTimeInMs uint, limit int) *responses.Response
}
