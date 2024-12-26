package repositories

import (
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

func (r *CategoryRepository) Create(category models.CategoryModel) string {
	utils.PanicIfError(r.db.Create(&category).Error)

	return category.Id
}

func (r *CategoryRepository) GetAll(limit, offset int) (categories []models.CategoryModel) {
	utils.PanicIfError(r.db.Limit(limit).Offset(offset).Preload("Products.Images").Find(&categories).Error)
	return
}

func (r *CategoryRepository) GetByName(name string) *models.CategoryModel {
	var category models.CategoryModel

	err := r.db.Where("name = ?", name).Preload("Products.Images").First(&category).Error

	return utils.HandleRecordNotFound[*models.CategoryModel](&category, err)
}

func (r *CategoryRepository) GetById(id string) *models.CategoryModel {
	var category models.CategoryModel
	err := r.db.Clauses(clause.Returning{}).Preload("Products.Images").Where("id = ?", id).First(&category).Error

	return utils.HandleRecordNotFound[*models.CategoryModel](&category, err)
}

func (r *CategoryRepository) UpdateById(id string, category models.CategoryModel) {
	utils.PanicIfError(r.db.Model(&models.CategoryModel{}).Where("id = ?", id).Updates(&category).Error)
}

func (r *CategoryRepository) DeleteById(id string) {
	utils.PanicIfError(r.db.Where("id = ?", id).Delete(&models.CategoryModel{}).Error)
}
