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

func (r *ProductRepository) GetAll(limit, offset int) ([]models.ProductModel, error) {
	var products []models.ProductModel
	err := r.db.Limit(limit).Offset(offset).Preload("Images").Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) Create(product models.ProductModel) (string, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return "", err
	}

	return product.Id, nil
}

func (r *ProductRepository) GetByName(name string) (*models.ProductModel, error) {
	var product models.ProductModel

	err := r.db.Where("name = ?", name).First(&product).Error

	return utils.HandleRecordNotFound[*models.ProductModel](&product, err)
}

func (r *ProductRepository) GetById(id string) (*models.ProductModel, error) {
	var product models.ProductModel
	err := r.db.Clauses(clause.Returning{}).Preload("Images").Where("id = ?", id).First(&product).Error

	return utils.HandleRecordNotFound[*models.ProductModel](&product, err)
}

func (r *ProductRepository) UpdateById(id string, product models.ProductModel) error {
	return r.db.Model(&models.ProductModel{}).Where("id = ?", id).Updates(&product).Error
}

func (r *ProductRepository) DeleteById(id string) error {
	return r.db.Where("id = ?", id).Delete(&models.ProductModel{}).Error
}

func (r *ProductRepository) GetByIds(ids []string) ([]models.ProductModel, error) {
	var products []models.ProductModel
	err := r.db.Where("id IN (?)", ids).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}
