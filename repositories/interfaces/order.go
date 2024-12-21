package interfaces

import "sushi-backend/models"

type IOrderRepository interface {
	GetAll(limit, offset int) ([]models.OrderModel, error)
	GetById(id uint) (*models.OrderModel, error)
	Create(category models.OrderModel) (uint, error)
	DeleteById(id uint) error
}
