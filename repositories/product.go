package repositories

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sushi-backend/models"
	"sushi-backend/repositories/dependencies"
	"sushi-backend/utils"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(deps dependencies.ProductRepositoryDependencies) *ProductRepository {
	if deps.Config.RunMigration() {
		utils.PanicIfError(deps.DB.AutoMigrate(&models.ProductModel{}))
	}

	return &ProductRepository{
		db: deps.DB,
	}
}

func (r *ProductRepository) GetAll(ctx context.Context, limit, offset int) (products []models.ProductModel) {
	utils.PanicIfErrorIsNotContextError(r.db.WithContext(ctx).Limit(limit).Offset(offset).Preload("Images").Find(&products).Error)

	return
}

func (r *ProductRepository) Create(ctx context.Context, product models.ProductModel) string {
	utils.PanicIfErrorIsNotContextError(r.db.WithContext(ctx).Create(&product).Error)

	return product.Id
}

func (r *ProductRepository) GetByName(ctx context.Context, name string) *models.ProductModel {
	var product models.ProductModel

	err := r.db.WithContext(ctx).Where("name = ?", name).First(&product).Error

	return utils.HandleRecordNotFound[*models.ProductModel](&product, err)
}

func (r *ProductRepository) GetById(ctx context.Context, id string) *models.ProductModel {
	var product models.ProductModel
	err := r.db.WithContext(ctx).Clauses(clause.Returning{}).Preload("Images").Where("id = ?", id).First(&product).Error

	return utils.HandleRecordNotFound[*models.ProductModel](&product, err)
}

func (r *ProductRepository) UpdateById(ctx context.Context, id string, product models.ProductModel) {
	utils.PanicIfErrorIsNotContextError(r.db.WithContext(ctx).Model(&models.ProductModel{}).Where("id = ?", id).Updates(&product).Error)
}

func (r *ProductRepository) DeleteById(ctx context.Context, id string) {
	utils.PanicIfErrorIsNotContextError(r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.ProductModel{}).Error)
}

func (r *ProductRepository) GetByIds(ctx context.Context, ids []string) (products []models.ProductModel) {
	utils.PanicIfErrorIsNotContextError(r.db.WithContext(ctx).Where("id IN (?)", ids).Find(&products).Error)

	return
}
