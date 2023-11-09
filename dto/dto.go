package dto

import (
	postmodel "crud/internal/modules/post/postModel"
	usermodel "crud/internal/modules/user/userModel"
	userresponse "crud/internal/modules/user/userResponse"
	"crud/response"
	"crud/utils"
)

func DtOPostResponse(request postmodel.Posts) response.PostResponse {
	return response.PostResponse{
		UserId:        request.UserId,
		Title:         request.Title,
		Body:          request.Body,
		CreatedAt:     request.CreatedAt,
		Category_Type: request.CategoryType,
	}
}

func ToDtoUser(user usermodel.Users) *userresponse.UserResponse {
	return &userresponse.UserResponse{
		UserID:    user.UserID,
		CreatedAt: utils.Now(),
		Email:     user.Email,
		Password:  user.Password,
	}
}
