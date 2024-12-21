package repositories

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sushi-backend/models"
	"sushi-backend/repositories/dependencies"
	"sushi-backend/utils"
)

type ProductImageRepository struct {
	db *gorm.DB
}

func NewProductImageRepository(deps dependencies.ProductImageRepositoryDependencies) *ProductImageRepository {
	if deps.Config.RunMigration() {
		utils.PanicIfError(deps.DB.AutoMigrate(&models.ProductImageModel{}))
	}

	return &ProductImageRepository{
		db: deps.DB,
	}
}

func (r *ProductImageRepository) Create(image models.ProductImageModel) (string, error) {
	err := r.db.Create(&image).Error
	if err != nil {
		return "", err
	}

	return image.Id, nil
}

func (r *ProductImageRepository) GetById(id string) (*models.ProductImageModel, error) {
	var image models.ProductImageModel
	err := r.db.Clauses(clause.Returning{}).Where("id = ?", id).First(&image).Error

	return utils.HandleRecordNotFound[*models.ProductImageModel](&image, err)
}

func (r *ProductImageRepository) DeleteById(id string) error {
	return r.db.Where("id = ?", id).Delete(&models.ProductImageModel{}).Error
}
