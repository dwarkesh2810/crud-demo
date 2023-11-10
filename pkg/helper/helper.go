package helper

import (
	userresponse "crud/internal/modules/user/userResponse"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Success int         `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func HashData(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

func VerifyHashedData(hashedString string, dataString string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(dataString), []byte(hashedString))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("email or password is incorrect")
		check = false
		return check, msg
	}
	return check, msg
}

func GetNewUUID() string {
	newUUID := uuid.New()
	return newUUID.String()
}

func GetTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")

	// Check if the header is present and starts with "Bearer "
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", fmt.Errorf("Invalid or Missing Token")
	}

	// Extract the token without the "Bearer " prefix
	authToken := authHeader[7:]

	return authToken, nil

}

func ParseError(err error) []string {
	data := strings.Split(err.Error(), "Error:")
	return data
}

func DataMarshal(data interface{}) (interface{}, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func JsonResponse(c *gin.Context, statusCode int, success int, data interface{}, err string) {
	var response Response = Response{
		Success: success,
		Data:    data,
		Error:   err,
	}
	c.JSON(statusCode, response)
}

func GetDataFromContext(c *gin.Context, key string) string {
	var userID string
	data, ok := c.Get(key)

	if ok {
		if data, ok := data.(userresponse.UserResponse); ok {
			userID = data.UserId
		}
	} else {

		return ""
	}

	log.Print(userID)
	return userID
}
