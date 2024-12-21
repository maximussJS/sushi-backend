package interfaces

import "sushi-backend/types/responses"

type IOrderFlowService interface {
	StartProcessing(id uint) *responses.Response
	ReadyToDeliver(id uint) *responses.Response
	StartDelivering(id uint) *responses.Response
	Delivered(id uint) *responses.Response
	Cancel(id uint) *responses.Response
}
