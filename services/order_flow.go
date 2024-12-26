package services

import (
	"fmt"
	"sushi-backend/config"
	"sushi-backend/constants"
	"sushi-backend/internal/telegram"
	"sushi-backend/models"
	"sushi-backend/repositories/interfaces"
	dependencies "sushi-backend/services/dependecies"
	"sushi-backend/types/responses"
	"time"
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

func (o *OrderFlowService) StartProcessing(id, estimatedTimeInMs uint) *responses.Response {
	order := o.orderRepository.GetById(id)

	if order == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Order with id %d not found", id))
	}

	if order.Status != constants.StatusCreated {
		return responses.NewBadRequestResponse(fmt.Sprintf("Order with id %d is not in created status", id))
	}

	o.orderRepository.UpdateById(id, models.OrderModel{
		Status: constants.StatusInProgress,
	})

	estimatedTime := time.Now().Add(time.Duration(estimatedTimeInMs) * time.Millisecond).Format("15:04:05")

	msg := fmt.Sprintf("Order with id %d is processing. Please be ready to take the order at %s", id, estimatedTime)

	o.telegram.SendMessageToChannel(o.config.TelegramDeliveryChatId(), msg, true)

	return responses.NewSuccessResponse(nil)
}

func (o *OrderFlowService) ReadyToDeliver(id uint) *responses.Response {
	order := o.orderRepository.GetById(id)

	if order == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Order with id %d not found", id))
	}

	if order.Status != constants.StatusInProgress {
		return responses.NewBadRequestResponse(fmt.Sprintf("Order with id %d is not in processing status", id))
	}

	o.orderRepository.UpdateById(id, models.OrderModel{
		Status:      constants.StatusReadyToDelivery,
		ProcessedAt: time.Now(),
	})

	o.telegram.SendMessageToChannel(o.config.TelegramDeliveryChatId(), fmt.Sprintf("Order with id %d is ready to deliver", id), true)

	return responses.NewSuccessResponse(nil)
}

func (o *OrderFlowService) StartDelivering(id, estimatedTimeInMs uint) *responses.Response {
	order := o.orderRepository.GetById(id)

	if order == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Order with id %d not found", id))
	}

	if order.Status != constants.StatusReadyToDelivery {
		return responses.NewBadRequestResponse(fmt.Sprintf("Order with id %d is not in ready to deliver status", id))
	}

	o.orderRepository.UpdateById(id, models.OrderModel{
		Status: constants.StatusInDelivery,
	})

	deliveryTime := time.Now().Add(time.Duration(estimatedTimeInMs) * time.Millisecond).Format("15:04:05")

	msg := fmt.Sprintf("Order with id %d is delivering. Approximately delivery time  %s", id, deliveryTime)

	o.telegram.SendMessageToChannel(o.config.TelegramDeliveryChatId(), msg, true)

	return responses.NewSuccessResponse(nil)
}

func (o *OrderFlowService) Delivered(id uint) *responses.Response {
	order := o.orderRepository.GetById(id)

	if order == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Order with id %d not found", id))
	}

	if order.Status != constants.StatusInDelivery {
		return responses.NewBadRequestResponse(fmt.Sprintf("Order with id %d is not in delivering status", id))
	}

	o.orderRepository.UpdateById(id, models.OrderModel{
		Status:      constants.StatusDelivered,
		DeliveredAt: time.Now(),
	})

	o.telegram.SendMessageToChannel(o.config.TelegramOrdersChatId(), fmt.Sprintf("Order with id %d is delivered", id), true)

	return responses.NewSuccessResponse(nil)
}

func (o *OrderFlowService) Cancel(id uint) *responses.Response {
	order := o.orderRepository.GetById(id)

	if order == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Order with id %d not found", id))
	}

	o.orderRepository.UpdateById(id, models.OrderModel{
		Status: constants.StatusCanceled,
	})

	o.telegram.SendMessageToChannel(o.config.TelegramOrdersChatId(), fmt.Sprintf("Order with id %d is canceled", id), true)

	return responses.NewSuccessResponse(nil)
}
