package dto

import (
	"crud/models"
	"crud/response"
	"crud/utils"
)

func DtOPostResponse(request models.Posts) response.PostResponse {
	return response.PostResponse{
		UserId:        request.UserId,
		Title:         request.Title,
		Body:          request.Body,
		CreatedAt:     request.CreatedAt,
		Category_Type: request.CategoryType,
	}
}

func ToDtoUser(user models.Users) *response.UserDataResponse {
	return &response.UserDataResponse{
		UserID:    user.UserID,
		CreatedAt: utils.Now(),
		Email:     user.Email,
		Password:  user.Password,
	}
}
