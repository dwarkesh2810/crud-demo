package postrepository

import (
	"crud/initialise"

	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func New() *PostRepository {
	return &PostRepository{
		DB: initialise.DB,
	}
}
