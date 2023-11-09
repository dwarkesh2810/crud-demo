package categoryrepository

import (
	"crud/pkg/config"

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

func (categoryRepository *CategoryRepository) List() {
	
} 