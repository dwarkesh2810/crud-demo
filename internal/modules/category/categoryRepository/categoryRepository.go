package categoryrepository

import (
	"crud/initialise"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func New() *CategoryRepository {
	return &CategoryRepository{
		DB: initialise.DB,
	}
}
