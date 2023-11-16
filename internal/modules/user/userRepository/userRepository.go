package userrepository

import (
	usermodel "crud/internal/modules/user/userModel"
	"crud/pkg/config"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func New() *UserRepository {
	return &UserRepository{
		DB: config.Connection(),
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

func (userRepository *UserRepository) FindByID(userId uint) usermodel.Users {
	var user usermodel.Users
	userRepository.DB.First(&user, "user_id= ?", userId)
	return user
}

func (userRepository *UserRepository) LastInsertedUserId() uint {
	var user usermodel.Users
	userRepository.DB.Last(&user)

	return user.UserId
}
