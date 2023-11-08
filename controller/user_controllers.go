package controller

import (
	"crud/initialise"
	"crud/models"
	"crud/request"
	"crud/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Sugnup(c *gin.Context) {
	// Get data from request body
	var body request.RegisterRequest

	err := c.Bind(&body)

	if err != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, "", utils.ParseError(err)[1])
		return
	}

	hashedPassword, err := utils.HashData(body.Password)

	if err != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, "", "Failed to hash password")
	}

	user := models.Users{
		UserID:    utils.GetNewUUID(),
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  string(hashedPassword),
		CreatedAt: utils.Now(),
		UserRole:  "user",
	}
	result := initialise.DB.Create(&user)

	if result.Error != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, "", "Unexpected database error")
		return
	}

	utils.JsonResponse(c, http.StatusOK, 1, user, "")
}

func Login(c *gin.Context) {
	// Get data from body
	var body request.LoginRequest
	var user models.Users

	err := c.ShouldBindJSON(&body)

	if err != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, "", utils.ParseError(err)[1])
		return
	}

	result := initialise.DB.Table("users").Select("id", "user_id", "email", "password").Where("email = ?", body.Email).Scan(&user)

	if result.Error != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, "", "Unexpected database error")
		return
	}

	if user.ID == 0 || user.UserID == "" {
		utils.JsonResponse(c, http.StatusBadRequest, 0, "", "Invalid Email")
		return
	}

	ok, errors := utils.VerifyHashedData(body.Password, user.Password)

	if !ok {
		utils.JsonResponse(c, http.StatusBadRequest, 0, "", errors)
		return
	}

	// Generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.UserID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	accessToken, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, "", "failed to create Token")
		return
	}

	utils.JsonResponse(c, http.StatusOK, 1, accessToken, "")
}
