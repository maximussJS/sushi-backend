package repositories

import (
	"context"
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

func (r *ProductImageRepository) Create(ctx context.Context, image models.ProductImageModel) string {
	utils.PanicIfErrorIsNotContextError(r.db.WithContext(ctx).Create(&image).Error)

	return image.Id
}

func (r *ProductImageRepository) GetById(ctx context.Context, id string) *models.ProductImageModel {
	var image models.ProductImageModel
	err := r.db.WithContext(ctx).Clauses(clause.Returning{}).Where("id = ?", id).First(&image).Error

	return utils.HandleRecordNotFound[*models.ProductImageModel](&image, err)
}

func (r *ProductImageRepository) DeleteById(ctx context.Context, id string) {
	utils.PanicIfErrorIsNotContextError(r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.ProductImageModel{}).Error)
}
