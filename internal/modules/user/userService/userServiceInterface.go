package userservice

import (
	userrequest "crud/internal/modules/user/userRequest"
	userresponse "crud/internal/modules/user/userResponse"
)

type UserServiceInterface interface {
	Create(request userrequest.RegisterRequest) (userresponse.UserResponse, error)
	CheckUserExist(email string) bool
	Login(userrequest.LoginRequest) (userresponse.LoginUserResponse, error)
}
