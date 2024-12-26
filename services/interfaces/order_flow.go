package interfaces

import (
	"context"
	"sushi-backend/types/responses"
)

type IOrderFlowService interface {
	StartProcessing(ctx context.Context, id, estimatedTimeInMs uint) *responses.Response
	ReadyToDeliver(ctx context.Context, id uint) *responses.Response
	StartDelivering(ctx context.Context, id, estimatedTimeInMs uint) *responses.Response
	Delivered(ctx context.Context, id uint) *responses.Response
	Cancel(ctx context.Context, id uint) *responses.Response
}
