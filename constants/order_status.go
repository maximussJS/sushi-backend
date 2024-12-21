package constants

type OrderStatus string

const (
	StatusCreated         OrderStatus = "created"
	StatusInProgress      OrderStatus = "in_progress"
	StatusReadyToDelivery OrderStatus = "ready_to_delivery"
	StatusInDelivery      OrderStatus = "in_delivery"
	StatusDelivered       OrderStatus = "delivered"
	StatusCanceled        OrderStatus = "canceled"
)
