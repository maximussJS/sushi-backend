package interfaces

import (
	"context"
	"sushi-backend/types/responses"
)

type IAnalyticService interface {
	GetOrdersAnalytic(ctx context.Context, startTimeInMs uint) *responses.Response
	GetTopOrderedProducts(ctx context.Context, startTimeInMs uint, limit int) *responses.Response
}
