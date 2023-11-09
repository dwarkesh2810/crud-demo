package postcontroller

import (
	postrequest "crud/internal/modules/post/postRequest"
	postservice "crud/internal/modules/post/postService"
	"crud/utils"
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

	utils.JsonResponse(c, http.StatusOK, 1, posts, "")
}

func (postcontrollers *PostControllers) CreatePost(c *gin.Context) {
	var request postrequest.PostCreateRequest
	userID := utils.GetDataFromContext(c, "user")

	if userID == "" {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "can not get data from context")
	}

	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Print("1")
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, utils.ParseError(err)[1])
		return
	}
	log.Print(request)
	post, err := postcontrollers.postService.StorePost(userID, request)

	if err != nil {
		log.Print("2")
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, err.Error())
		return
	}

	utils.JsonResponse(c, http.StatusOK, 1, post, "")
}
