package repositories

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sushi-backend/constants"
	"sushi-backend/models"
	"sushi-backend/repositories/dependencies"
	"sushi-backend/utils"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(deps dependencies.OrderRepositoryDependencies) *OrderRepository {
	if deps.Config.RunMigration() {
		utils.PanicIfError(deps.DB.AutoMigrate(&models.OrderProductModel{}))
		utils.PanicIfError(deps.DB.AutoMigrate(&models.OrderModel{}))
	}

	return &OrderRepository{
		db: deps.DB,
	}
}

func (r *OrderRepository) Create(order models.OrderModel) (uint, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return 0, err
	}

	return order.Id, nil
}

func (r *OrderRepository) GetAll(limit, offset int) ([]models.OrderModel, error) {
	var orders []models.OrderModel
	err := r.db.Limit(limit).Offset(offset).Preload("OrderedProducts.Product").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) GetById(id uint) (*models.OrderModel, error) {
	var order models.OrderModel
	err := r.db.Clauses(clause.Returning{}).Preload("OrderedProducts.Product").Where("id = ?", id).First(&order).Error

	return utils.HandleRecordNotFound[*models.OrderModel](&order, err)
}

func (r *OrderRepository) DeleteById(id uint) error {
	return r.db.Where("id = ?", id).Delete(&models.OrderModel{}).Error
}

func (r *OrderRepository) UpdateStatusById(id uint, status constants.OrderStatus) error {
	return r.db.Model(&models.OrderModel{}).Where("id = ?", id).Update("status", status).Error
}
