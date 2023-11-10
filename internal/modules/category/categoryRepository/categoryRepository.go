package categoryrepository

import (
	categorymodel "crud/internal/modules/category/categoryModel"
	"crud/pkg/config"
	"errors"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func New() *CategoryRepository {
	return &CategoryRepository{
		DB: config.Connection(),
	}
}

func (categoryRepository *CategoryRepository) List() []categorymodel.Category {
	var categories []categorymodel.Category

	categoryRepository.DB.Order("created_at desc").Find(&categories)

	return categories
}

func (categoryRepository *CategoryRepository) FindByID(id int) (categorymodel.Category, error) {
	var category categorymodel.Category

	result := config.DB.First(&category, "id = ?", id)

	if result.Error != nil {
		return category, errors.New("unexpected Database Error")
	}
	return category, nil
}

func (categoryRepository *CategoryRepository) Create(category categorymodel.Category) (categorymodel.Category, error) {
	var newCategory categorymodel.Category

	result := categoryRepository.DB.Create(&category).Scan(&newCategory)
	if result.Error != nil {
		return newCategory, errors.New("unexpected Database Error")
	}

	return newCategory, nil
}

func (categoryRepository *CategoryRepository) FindByCategory(category string) (categorymodel.Category, error) {
	var cat categorymodel.Category

	result := config.DB.First(&cat, "category_type = ?", category)

	if result.Error != nil {
		return cat, errors.New("unexpected Database Error")
	}
	return cat, nil
}
