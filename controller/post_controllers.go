package controller

import (
	"crud/dto"
	"crud/initialise"
	"crud/models"
	"crud/request"
	"crud/response"
	"crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	var posts []models.Posts

	result := initialise.DB.Find(&posts)

	if result.Error != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "Unexpected database error")
		return
	}

	utils.JsonResponse(c, http.StatusOK, 1, posts, "")
}

func CreatePost(c *gin.Context) {
	var p models.Posts
	var post request.PostCreateRequest
	var userID string
	user, ok := c.Get("user")

	if ok {
		if user, ok := user.(*response.UserDataResponse); ok {
			userID = user.UserID
		} else {
			utils.JsonResponse(c, http.StatusUnauthorized, 0, nil, "")
			return
		}
	} else {
		utils.JsonResponse(c, http.StatusUnauthorized, 0, nil, "")
		return
	}

	err := c.ShouldBindJSON(&post)
	if err != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, utils.ParseError(err)[1])
		return
	}

	// p = models.Posts{
	// 	Body:         post.Body,
	// 	Title:        post.Title,
	// 	UserId:       userID,
	// 	CreatedAt:    utils.Now(),
	// 	CategoryType: post.Category,
	// }

	p.Body = post.Body
	p.Title = post.Title
	p.UserId = userID
	p.CreatedAt = utils.Now()
	p.CategoryType = post.Category

	result := initialise.DB.Create(&p)

	if result.Error != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "Unexpected database error")
		return
	}

	utils.JsonResponse(c, http.StatusOK, 1, dto.DtOPostResponse(p), "")
}

func UpdatePost(c *gin.Context) {
	var post models.Posts
	var updatedPost request.PostUpdateRequest
	var userID string

	user, ok := c.Get("user")

	if ok {
		if user, ok := user.(*response.UserDataResponse); ok {
			userID = user.UserID
		} else {
			utils.JsonResponse(c, http.StatusUnauthorized, 0, nil, "")
			return
		}
	} else {
		utils.JsonResponse(c, http.StatusUnauthorized, 0, nil, "")
		return
	}

	err := c.ShouldBindJSON(&updatedPost)

	if err != nil {
		utils.JsonResponse(c, http.StatusUnauthorized, 0, nil, utils.ParseError(err)[1])
		return
	}

	result := initialise.DB.First(&post, "id = ?", updatedPost.ID)

	if result.Error != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "Unexpected database error")
		return
	}

	if !(post.UserId == userID) {
		utils.JsonResponse(c, http.StatusUnauthorized, 0, nil, "you can not modify others posts")
		return
	}

	post.Body = updatedPost.Body
	post.Title = updatedPost.Title
	post.CategoryType = updatedPost.Category

	result = initialise.DB.Save(&post)

	if result.Error != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "Unexpected database error")
		return
	}

	utils.JsonResponse(c, http.StatusOK, 1, dto.DtOPostResponse(post), "")
}

func DeletePost(c *gin.Context) {
	var post models.Posts
	var userID string
	var deletePost request.PostDeleteRequest
	user, ok := c.Get("user")

	if ok {
		if user, ok := user.(*response.UserDataResponse); ok {
			userID = user.UserID
		} else {
			utils.JsonResponse(c, http.StatusUnauthorized, 0, nil, "")
			return
		}
	} else {
		utils.JsonResponse(c, http.StatusUnauthorized, 0, nil, "")
		return
	}

	err := c.ShouldBindJSON(&deletePost)

	if err != nil {
		utils.JsonResponse(c, http.StatusUnauthorized, 0, nil, utils.ParseError(err)[1])
		return
	}

	result := initialise.DB.First(&post, "id = ?", deletePost.ID)

	if result.Error != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "Unexpected database error")
		return
	}

	if !(post.UserId == userID) {
		utils.JsonResponse(c, http.StatusUnauthorized, 0, nil, "you can not Delete others posts")
		return
	}

	result = initialise.DB.Delete(&post)

	if result.Error != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "Unexpected database error")
		return
	}

	utils.JsonResponse(c, http.StatusOK, 1, dto.DtOPostResponse(post), "")
}

func GetPostByCategory(c *gin.Context) {
	cat := c.Param("category")

	if cat == ""{
		
	}
}
