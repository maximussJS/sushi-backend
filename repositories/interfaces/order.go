package interfaces

import (
	"context"
	"sushi-backend/models"
	"sushi-backend/types/analytic"
	"time"
)

type IOrderRepository interface {
	GetAll(ctx context.Context, limit, offset int) []models.OrderModel
	GetById(ctx context.Context, id uint) *models.OrderModel
	Create(ctx context.Context, category models.OrderModel) uint
	DeleteById(ctx context.Context, id uint)
	UpdateById(ctx context.Context, id uint, order models.OrderModel)
	GetDeliveredOrdersAnalytic(ctx context.Context, startTime time.Time) analytic.OrderAnalytic
	GetTopOrderedProducts(ctx context.Context, startTime time.Time, limit int) []analytic.TopProduct
}
