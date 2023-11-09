package userrepository

import (
	usermodel "crud/internal/modules/user/userModel"
	"crud/pkg/database"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func New() *UserRepository {
	return &UserRepository{
		DB: database.Connection(),
	}
}

func (userRepository *UserRepository) Create(user usermodel.Users) (usermodel.Users, error) {
	var newUser usermodel.Users
	result := userRepository.DB.Create(&user).Scan(&newUser)

	if result.Error != nil {
		return newUser, errors.New("Unexpected database error")
	}
	return newUser, nil
}

func (userRepository *UserRepository) FindByEmail(email string) usermodel.Users {
	var user usermodel.Users
	userRepository.DB.First(&user, "email= ?", email)
	return user
}

func (userRepository *UserRepository) FindByID(id int) usermodel.Users {
	var user usermodel.Users
	userRepository.DB.First(&user, "id= ?", id)
	return user
}
