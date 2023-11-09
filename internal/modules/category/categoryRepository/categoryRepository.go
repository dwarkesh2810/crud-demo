package categoryrepository

import (
	"crud/pkg/database"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func New() *CategoryRepository {
	return &CategoryRepository{
		DB: database.Connection(),
	}
}
