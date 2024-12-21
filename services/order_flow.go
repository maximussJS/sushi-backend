package services

import (
	"fmt"
	"sushi-backend/config"
	"sushi-backend/constants"
	"sushi-backend/internal/telegram"
	"sushi-backend/repositories/interfaces"
	dependencies "sushi-backend/services/dependecies"
	"sushi-backend/types/responses"
	"sushi-backend/utils"
)

type OrderFlowService struct {
	config          config.IConfig
	telegram        telegram.ITelegram
	orderRepository interfaces.IOrderRepository
}

func NewOrderFlowService(deps dependencies.OrderFlowServiceDependencies) *OrderFlowService {
	return &OrderFlowService{
		config:          deps.Config,
		telegram:        deps.Telegram,
		orderRepository: deps.OrderRepository,
	}
}

/*
StartProcessing(id uint)
	ReadyToDeliver(id uint)
	StartDelivering(id uint)
	Delivered(id uint)
	Cancel(id uint)
*/

func (o *OrderFlowService) StartProcessing(id uint) *responses.Response {
	order := utils.PanicIfErrorWithResultReturning(o.orderRepository.GetById(id))

	if order == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Order with id %d not found", id))
	}

	if order.Status != constants.StatusCreated {
		return responses.NewBadRequestResponse(fmt.Sprintf("Order with id %d is not in created status", id))
	}

	utils.PanicIfError(o.orderRepository.UpdateStatusById(id, constants.StatusInProgress))

	return responses.NewSuccessResponse(nil)
}

func (o *OrderFlowService) ReadyToDeliver(id uint) *responses.Response {
	order := utils.PanicIfErrorWithResultReturning(o.orderRepository.GetById(id))

	if order == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Order with id %d not found", id))
	}

	if order.Status != constants.StatusInProgress {
		return responses.NewBadRequestResponse(fmt.Sprintf("Order with id %d is not in processing status", id))
	}

	utils.PanicIfError(o.orderRepository.UpdateStatusById(id, constants.StatusReadyToDelivery))

	o.telegram.SendMessageToChannel(o.config.TelegramDeliveryChatId(), fmt.Sprintf("Order with id %d is ready to deliver", id), true)

	return responses.NewSuccessResponse(nil)
}

func (o *OrderFlowService) StartDelivering(id uint) *responses.Response {
	order := utils.PanicIfErrorWithResultReturning(o.orderRepository.GetById(id))

	if order == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Order with id %d not found", id))
	}

	if order.Status != constants.StatusReadyToDelivery {
		return responses.NewBadRequestResponse(fmt.Sprintf("Order with id %d is not in ready to deliver status", id))
	}

	utils.PanicIfError(o.orderRepository.UpdateStatusById(id, constants.StatusInDelivery))

	return responses.NewSuccessResponse(nil)
}

func (o *OrderFlowService) Delivered(id uint) *responses.Response {
	order := utils.PanicIfErrorWithResultReturning(o.orderRepository.GetById(id))

	if order == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Order with id %d not found", id))
	}

	if order.Status != constants.StatusInDelivery {
		return responses.NewBadRequestResponse(fmt.Sprintf("Order with id %d is not in delivering status", id))
	}

	utils.PanicIfError(o.orderRepository.UpdateStatusById(id, constants.StatusDelivered))

	o.telegram.SendMessageToChannel(o.config.TelegramOrdersChatId(), fmt.Sprintf("Order with id %d is delivered", id), true)

	return responses.NewSuccessResponse(nil)
}

func (o *OrderFlowService) Cancel(id uint) *responses.Response {
	order := utils.PanicIfErrorWithResultReturning(o.orderRepository.GetById(id))

	if order == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Order with id %d not found", id))
	}

	utils.PanicIfError(o.orderRepository.UpdateStatusById(id, constants.StatusCanceled))

	o.telegram.SendMessageToChannel(o.config.TelegramOrdersChatId(), fmt.Sprintf("Order with id %d is canceled", id), true)

	return responses.NewSuccessResponse(nil)
}
