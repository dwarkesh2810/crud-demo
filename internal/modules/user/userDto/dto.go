package userdto

import (
	usermodel "crud/internal/modules/user/userModel"
	userresponse "crud/internal/modules/user/userResponse"
	"crud/pkg/helper"
)

func ToUser(user usermodel.Users) userresponse.UserResponse {
	return userresponse.UserResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserId:    user.UserId,
		CreatedAt: helper.Now(),
		Email:     user.Email,
		Password:  user.Password,
	}
}

func ToUserLogin(user usermodel.Users, token string) userresponse.LoginUserResponse {
	return userresponse.LoginUserResponse{
		UserResponse: ToUser(user),
		Token:        token,
	}
}
