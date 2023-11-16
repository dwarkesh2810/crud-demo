package userrepository

import (
	usermodel "crud/internal/modules/user/userModel"
)

type UserRepositoryInterface interface {
	Create(usermodel.Users) (usermodel.Users, error)
	FindByEmail(string) usermodel.Users
	FindByID(uint) usermodel.Users
	LastInsertedUserId() uint
}
