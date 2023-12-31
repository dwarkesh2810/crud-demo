package middleware

import (
	userdto "crud/internal/modules/user/userDto"
	userrepository "crud/internal/modules/user/userRepository"
	"crud/pkg/helper"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := helper.GetTokenFromHeader(c)

	if err != nil {
		helper.JsonResponse(c, http.StatusBadRequest, 0, nil, err.Error())
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check expiration time
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		userId := claims["sub"].(float64)
		userIdinInt := uint(userId)

		user := userrepository.New().FindByID(userIdinInt)

		if user.UserId == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// Attach to request
		users := userdto.ToUser(user)
		c.Set("user", users)

		// continue
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

}
