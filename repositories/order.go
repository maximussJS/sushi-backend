package repositories

import (
	"context"
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

func (r *OrderRepository) Create(ctx context.Context, order models.OrderModel) uint {
	utils.PanicIfErrorIsNotContextError(r.db.WithContext(ctx).Create(&order).Error)

	return order.Id
}

func (r *OrderRepository) GetAll(ctx context.Context, limit, offset int) (orders []models.OrderModel) {
	utils.PanicIfErrorIsNotContextError(r.db.WithContext(ctx).Limit(limit).Offset(offset).Preload("OrderedProducts.Product").Find(&orders).Error)

	return
}

func (r *OrderRepository) GetById(ctx context.Context, id uint) *models.OrderModel {
	var order models.OrderModel
	err := r.db.WithContext(ctx).Clauses(clause.Returning{}).Preload("OrderedProducts.Product").Where("id = ?", id).First(&order).Error

	return utils.HandleRecordNotFound[*models.OrderModel](&order, err)
}

func (r *OrderRepository) DeleteById(ctx context.Context, id uint) {
	utils.PanicIfErrorIsNotContextError(r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.OrderModel{}).Error)
}

func (r *OrderRepository) UpdateById(ctx context.Context, id uint, order models.OrderModel) {
	utils.PanicIfErrorIsNotContextError(r.db.WithContext(ctx).Model(&models.OrderModel{}).Where("id = ?", id).Updates(&order).Error)
}

func (r *OrderRepository) GetDeliveredOrdersAnalytic(ctx context.Context, startTime time.Time) (orderAnalytic analytic.OrderAnalytic) {
	utils.PanicIfErrorIsNotContextError(
		r.db.
			WithContext(ctx).
			Model(&models.OrderModel{}).
			Select("COUNT(*) as orders_count, COALESCE(SUM(price), 0) as total_amount").
			Where("status = ? AND created_at > ?", constants.StatusDelivered, startTime).
			Scan(&orderAnalytic).Error,
	)

	return
}

func (r *OrderRepository) GetTopOrderedProducts(ctx context.Context, startTime time.Time, limit int) (topProducts []analytic.TopProduct) {
	query := r.db.
		WithContext(ctx).
		Table("order_products").
		Select("products.id as product_id, products.name as product_name, SUM(order_products.quantity) as total_quantity").
		Joins("JOIN orders ON orders.id = order_products.order_id").
		Joins("JOIN products ON products.id = order_products.product_id").
		Where("orders.status = ? AND orders.created_at > ?", constants.StatusDelivered, startTime).
		Group("products.id, products.name").
		Order("total_quantity DESC").
		Limit(limit)

	utils.PanicIfErrorIsNotContextError(query.Scan(&topProducts).Error)

	return
}
