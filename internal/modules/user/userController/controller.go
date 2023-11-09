package usercontroller

import (
	userrequest "crud/internal/modules/user/userRequest"
	userservice "crud/internal/modules/user/userService"
	"crud/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService userservice.UserServiceInterface
}

func New() *UserController {
	return &UserController{
		userService: userservice.New(),
	}
}

func (userController *UserController) RegisterUser(c *gin.Context) {
	var request userrequest.RegisterRequest
	err := c.ShouldBindJSON(&request)

	if err != nil {
		helper.JsonResponse(c, http.StatusBadRequest, 0, nil, "Error While Binding Request")
		return
	}

	ok := userController.userService.CheckUserExist(request.Email)

	if ok {
		helper.JsonResponse(c, http.StatusConflict, 0, nil, "Email already registered")
		return
	}

	user, err := userController.userService.Create(request)

	if err != nil {
		helper.JsonResponse(c, http.StatusBadRequest, 0, nil, err.Error())
	}

	helper.JsonResponse(c, http.StatusCreated, 1, user, "")

}

func (userController *UserController) LoginUser(c *gin.Context) {
	var request userrequest.LoginRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		helper.JsonResponse(c, http.StatusBadRequest, 0, nil, "Error While Binding Request")
		return
	}

	data, err := userController.userService.Login(request)

	if err != nil {
		helper.JsonResponse(c, http.StatusBadRequest, 0, nil, "")
		return
	}

	helper.JsonResponse(c, http.StatusOK, 1, data, "")
}
