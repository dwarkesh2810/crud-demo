package postcontroller

import (
	postrequest "crud/internal/modules/post/postRequest"
	postservice "crud/internal/modules/post/postService"
	"crud/pkg/helper"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostControllers struct {
	postService postservice.PostServiceInterface
}

func New() *PostControllers {
	return &PostControllers{
		postService: postservice.New(),
	}
}

func (postControllers *PostControllers) ShowPost(c *gin.Context) {
	posts := postControllers.postService.GetPosts()

	helper.JsonResponse(c, http.StatusOK, 1, posts, "")
}

func (postcontrollers *PostControllers) CreatePost(c *gin.Context) {
	var request postrequest.PostCreateRequest
	userID := helper.GetDataFromContext(c, "user")

	if userID == "" {
		helper.JsonResponse(c, http.StatusBadRequest, 0, nil, "can not get data from context")
	}

	err := c.ShouldBindJSON(&request)
	if err != nil {
		helper.JsonResponse(c, http.StatusBadRequest, 0, nil, helper.ParseError(err)[1])
		return
	}

	post, err := postcontrollers.postService.CreatePost(userID, request)

	if err != nil {
		helper.JsonResponse(c, http.StatusBadRequest, 0, nil, err.Error())
		return
	}

	helper.JsonResponse(c, http.StatusOK, 1, post, "")
}

func (postControllers *PostControllers) UpdatePost(c *gin.Context) {
	var request postrequest.PostUpdateRequest
	userID := helper.GetDataFromContext(c, "user")

	if userID == "" {
		helper.JsonResponse(c, http.StatusBadRequest, 0, nil, "can not get data from context")
	}

	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Print("1")
		helper.JsonResponse(c, http.StatusBadRequest, 0, nil, helper.ParseError(err)[1])
		return
	}

	post, err := postControllers.postService.UpdatePost(userID, request)

	if err != nil {
		helper.JsonResponse(c, http.StatusBadRequest, 0, nil, err.Error())
		return
	}

	helper.JsonResponse(c, http.StatusOK, 1, post, "")
}
