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
	err := r.db.Limit(limit).Offset(offset).Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepository) FindByName(name string) (*models.CategoryModel, error) {
	var category models.CategoryModel

	err := r.db.Where("name = ?", name).First(&category).Error

	return utils.HandleRecordNotFound[*models.CategoryModel](&category, err)
}

func (r *CategoryRepository) FindById(id string) (*models.CategoryModel, error) {
	var category models.CategoryModel
	err := r.db.Clauses(clause.Returning{}).Where("id = ?", id).First(&category).Error

	return utils.HandleRecordNotFound[*models.CategoryModel](&category, err)
}
