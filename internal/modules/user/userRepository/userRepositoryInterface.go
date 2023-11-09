package userrepository

import usermodel "crud/internal/modules/user/userModel"

type UserRepositoryInterface interface {
	Create(usermodel.Users) (usermodel.Users, error)
	FindByEmail(email string) usermodel.Users
	FindByID(id int) usermodel.Users
}
