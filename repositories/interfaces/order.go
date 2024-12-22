package interfaces

import (
	"sushi-backend/models"
	"sushi-backend/types/analytic"
	"time"
)

type IOrderRepository interface {
	GetAll(limit, offset int) ([]models.OrderModel, error)
	GetById(id uint) (*models.OrderModel, error)
	Create(category models.OrderModel) (uint, error)
	DeleteById(id uint) error
	UpdateById(id uint, order models.OrderModel) error
	GetDeliveredOrdersAnalytic(startTime time.Time) (analytic.OrderAnalytic, error)
	GetTopOrderedProducts(startTime time.Time, limit int) ([]analytic.TopProduct, error)
}
