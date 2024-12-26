package repositories

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sushi-backend/models"
	"sushi-backend/repositories/dependencies"
	"sushi-backend/utils"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(deps dependencies.CategoryRepositoryDependencies) *CategoryRepository {
	if deps.Config.RunMigration() {
		utils.PanicIfError(deps.DB.AutoMigrate(&models.CategoryModel{}))
	}

	return &CategoryRepository{
		db: deps.DB,
	}
}

func (r *CategoryRepository) Create(ctx context.Context, category models.CategoryModel) string {
	utils.PanicIfErrorIsNotContextError(r.db.WithContext(ctx).Create(&category).Error)

	return category.Id
}

func (r *CategoryRepository) GetAll(ctx context.Context, limit, offset int) (categories []models.CategoryModel) {
	utils.PanicIfErrorIsNotContextError(r.db.WithContext(ctx).Limit(limit).Offset(offset).Preload("Products.Images").Find(&categories).Error)
	return
}

func (r *CategoryRepository) GetByName(ctx context.Context, name string) *models.CategoryModel {
	var category models.CategoryModel

	err := r.db.WithContext(ctx).Where("name = ?", name).Preload("Products.Images").First(&category).Error

	return utils.HandleRecordNotFound[*models.CategoryModel](&category, err)
}

func (r *CategoryRepository) GetById(ctx context.Context, id string) *models.CategoryModel {
	var category models.CategoryModel
	err := r.db.WithContext(ctx).Clauses(clause.Returning{}).Preload("Products.Images").Where("id = ?", id).First(&category).Error

	return utils.HandleRecordNotFound[*models.CategoryModel](&category, err)
}

func (r *CategoryRepository) UpdateById(ctx context.Context, id string, category models.CategoryModel) {
	utils.PanicIfErrorIsNotContextError(r.db.WithContext(ctx).Model(&models.CategoryModel{}).Where("id = ?", id).Updates(&category).Error)
}

func (r *CategoryRepository) DeleteById(ctx context.Context, id string) {
	utils.PanicIfErrorIsNotContextError(r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.CategoryModel{}).Error)
}
