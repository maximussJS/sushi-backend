package repositories

import (
	"gorm.io/gorm"
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

func (r *ProductRepository) GetAll() ([]models.ProductModel, error) {
	var products []models.ProductModel
	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}
