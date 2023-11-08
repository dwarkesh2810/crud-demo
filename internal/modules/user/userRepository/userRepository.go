package userrepository

import (
	"crud/initialise"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func New() *UserRepository {
	return &UserRepository{
		DB: initialise.DB,
	}
}
