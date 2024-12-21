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

func (r *CategoryRepository) Create(category models.CategoryModel) (string, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return "", err
	}

	return category.Id, nil
}

func (r *CategoryRepository) GetAll(limit, offset int) ([]models.CategoryModel, error) {
	var categories []models.CategoryModel
	err := r.db.Limit(limit).Offset(offset).Preload("Products.Images").Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepository) GetByName(name string) (*models.CategoryModel, error) {
	var category models.CategoryModel

	err := r.db.Where("name = ?", name).Preload("Products.Images").First(&category).Error

	return utils.HandleRecordNotFound[*models.CategoryModel](&category, err)
}

func (r *CategoryRepository) GetById(id string) (*models.CategoryModel, error) {
	var category models.CategoryModel
	err := r.db.Clauses(clause.Returning{}).Preload("Products.Images").Where("id = ?", id).First(&category).Error

	return utils.HandleRecordNotFound[*models.CategoryModel](&category, err)
}

func (r *CategoryRepository) UpdateById(id string, category models.CategoryModel) error {
	return r.db.Model(&models.CategoryModel{}).Where("id = ?", id).Updates(&category).Error
}

func (r *CategoryRepository) DeleteById(id string) error {
	return r.db.Where("id = ?", id).Delete(&models.CategoryModel{}).Error
}
