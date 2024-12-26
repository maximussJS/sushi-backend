package interfaces

import (
	"sushi-backend/models"
	"sushi-backend/types/analytic"
	"time"
)

type IOrderRepository interface {
	GetAll(limit, offset int) []models.OrderModel
	GetById(id uint) *models.OrderModel
	Create(category models.OrderModel) uint
	DeleteById(id uint)
	UpdateById(id uint, order models.OrderModel)
	GetDeliveredOrdersAnalytic(startTime time.Time) analytic.OrderAnalytic
	GetTopOrderedProducts(startTime time.Time, limit int) []analytic.TopProduct
}
