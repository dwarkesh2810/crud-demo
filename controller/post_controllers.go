package controller

import (
	"crud/dto"
	postmodel "crud/internal/modules/post/postModel"
	postservice "crud/internal/modules/post/postService"
	"crud/pkg/database"
	"crud/request"
	"crud/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controllers struct {
	postService postservice.PostServiceInterface
}

func New() *Controllers {
	return &Controllers{
		postService: postservice.New(),
	}
}

func GetPosts(c *gin.Context) {
	var posts []postmodel.Posts

	result := database.DB.Find(&posts)

	if result.Error != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "Unexpected database error")
		return
	}

	utils.JsonResponse(c, http.StatusOK, 1, posts, "")
}

func CreatePost(c *gin.Context) {
	var p postmodel.Posts
	var post request.PostCreateRequest
	userID := utils.GetDataFromContext(c, "user")

	if userID == "" {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "can not get data from context")
	}

	err := c.ShouldBindJSON(&post)
	if err != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, utils.ParseError(err)[1])
		return
	}
	log.Print(post)

	p.Body = post.Body
	p.Title = post.Title
	p.UserId = userID
	p.CreatedAt = utils.Now()
	p.CategoryType = post.Category

	result := database.DB.Create(&p)

	if result.Error != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "Unexpected database error")
		return
	}

	utils.JsonResponse(c, http.StatusOK, 1, dto.DtOPostResponse(p), "")
}

func UpdatePost(c *gin.Context) {
	var post postmodel.Posts
	var updatedPost request.PostUpdateRequest

	userID := utils.GetDataFromContext(c, "user")

	if userID == "" {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "can not get data from context")
		return
	}

	err := c.ShouldBindJSON(&updatedPost)

	if err != nil {
		utils.JsonResponse(c, http.StatusUnauthorized, 0, nil, utils.ParseError(err)[1])
		return
	}

	result := database.DB.First(&post, "id = ?", updatedPost.ID)

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
	result = database.DB.Save(&post)

	if result.Error != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "Unexpected database error")
		return
	}

	utils.JsonResponse(c, http.StatusOK, 1, dto.DtOPostResponse(post), "")
}

func DeletePost(c *gin.Context) {
	var post postmodel.Posts
	var deletePost request.PostDeleteRequest
	userID := utils.GetDataFromContext(c, "user")

	if userID == "" {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "can not get data from context")
		return
	}

	err := c.ShouldBindJSON(&deletePost)

	if err != nil {
		utils.JsonResponse(c, http.StatusUnauthorized, 0, nil, utils.ParseError(err)[1])
		return
	}

	result := database.DB.First(&post, "id = ?", deletePost.ID)

	if result.Error != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "Unexpected database error")
		return
	}

	if !(post.UserId == userID) {
		utils.JsonResponse(c, http.StatusUnauthorized, 0, nil, "you can not Delete others posts")
		return
	}

	result = database.DB.Delete(&post)

	if result.Error != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "Unexpected database error")
		return
	}

	utils.JsonResponse(c, http.StatusOK, 1, dto.DtOPostResponse(post), "")
}

func GetPostByCategory(c *gin.Context) {
	var posts []postmodel.Posts
	var cat request.CategoryRequest

	err := c.ShouldBindJSON(&cat)

	if err != nil {
		utils.JsonResponse(c, http.StatusUnauthorized, 0, nil, utils.ParseError(err)[1])
		return
	}

	result := database.DB.Find(&posts, "category_type = ?", cat.Category)

	if result.Error != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "Unexpected database error")
		return
	}

	utils.JsonResponse(c, http.StatusOK, 1, posts, "")
}

func GetPostByUser(c *gin.Context) {
	var posts []postmodel.Posts
	userID := utils.GetDataFromContext(c, "user")

	if userID == "" {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "can not get data from context")
		return
	}

	result := database.DB.Find(&posts, "user_id = ?", userID)
	if result.Error != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "Unexpected database error")
		return
	}
	utils.JsonResponse(c, http.StatusOK, 1, posts, "")

}

func GetPostDateByUserAndCategory(c *gin.Context) {
	var posts []postmodel.Posts
	var reqPost request.UserAndCategoryRequest

	err := c.ShouldBindJSON(&reqPost)

	if err != nil {
		utils.JsonResponse(c, http.StatusUnauthorized, 0, nil, utils.ParseError(err)[1])
		return
	}
	db := database.DB
	result := db.Where(
		db.Where("user_id = ?", reqPost.UserID).Where(db.Where("category_type = ?", reqPost.Category))).Find(&posts)

	if result.Error != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "Unexpected database error")
		return
	}

	utils.JsonResponse(c, http.StatusOK, 1, posts, "")
}
