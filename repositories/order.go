package repositories

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sushi-backend/constants"
	"sushi-backend/models"
	"sushi-backend/repositories/dependencies"
	"sushi-backend/types/analytic"
	"sushi-backend/utils"
	"time"
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

func (r *OrderRepository) UpdateById(id uint, order models.OrderModel) error {
	return r.db.Model(&models.OrderModel{}).Where("id = ?", id).Updates(&order).Error
}

func (r *OrderRepository) GetDeliveredOrdersAnalytic(startTime time.Time) (analytic.OrderAnalytic, error) {
	var orderAnalytic analytic.OrderAnalytic

	// Perform the query
	err := r.db.Model(&models.OrderModel{}).
		Select("COUNT(*) as orders_count, COALESCE(SUM(price), 0) as total_amount").
		Where("status = ? AND created_at > ?", constants.StatusDelivered, startTime).
		Scan(&orderAnalytic).Error

	if err != nil {
		return orderAnalytic, err
	}

	return orderAnalytic, nil
}

func (r *OrderRepository) GetTopOrderedProducts(startTime time.Time, limit int) ([]analytic.TopProduct, error) {
	var topProducts []analytic.TopProduct

	query := r.db.Table("order_products").
		Select("products.id as product_id, products.name as product_name, SUM(order_products.quantity) as total_quantity").
		Joins("JOIN orders ON orders.id = order_products.order_id").
		Joins("JOIN products ON products.id = order_products.product_id").
		Where("orders.status = ? AND orders.created_at > ?", constants.StatusDelivered, startTime).
		Group("products.id, products.name").
		Order("total_quantity DESC").
		Limit(limit)

	err := query.Scan(&topProducts).Error
	if err != nil {
		return nil, err
	}

	return topProducts, nil
}
