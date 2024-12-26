package repositories

import (
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

func (r *ProductRepository) GetAll(limit, offset int) (products []models.ProductModel) {
	utils.PanicIfError(r.db.Limit(limit).Offset(offset).Preload("Images").Find(&products).Error)

	return
}

func (r *ProductRepository) Create(product models.ProductModel) string {
	utils.PanicIfError(r.db.Create(&product).Error)

	return product.Id
}

func (r *ProductRepository) GetByName(name string) *models.ProductModel {
	var product models.ProductModel

	err := r.db.Where("name = ?", name).First(&product).Error

	return utils.HandleRecordNotFound[*models.ProductModel](&product, err)
}

func (r *ProductRepository) GetById(id string) *models.ProductModel {
	var product models.ProductModel
	err := r.db.Clauses(clause.Returning{}).Preload("Images").Where("id = ?", id).First(&product).Error

	return utils.HandleRecordNotFound[*models.ProductModel](&product, err)
}

func (r *ProductRepository) UpdateById(id string, product models.ProductModel) {
	utils.PanicIfError(r.db.Model(&models.ProductModel{}).Where("id = ?", id).Updates(&product).Error)
}

func (r *ProductRepository) DeleteById(id string) {
	utils.PanicIfError(r.db.Where("id = ?", id).Delete(&models.ProductModel{}).Error)
}

func (r *ProductRepository) GetByIds(ids []string) (products []models.ProductModel) {
	utils.PanicIfError(r.db.Where("id IN (?)", ids).Find(&products).Error)

	return
}
