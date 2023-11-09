package middleware

import (
	"crud/dto"
	usermodel "crud/internal/modules/user/userModel"
	"crud/pkg/database"
	"crud/utils"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := utils.GetTokenFromHeader(c)

	if err != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, err.Error())
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
		}
		// Find the user with token sub
		var user usermodel.Users
		// initializers.DB.Table("users").Select("id", "email", "password", ).Where("ID = ?", claims["sub"]).Scan(&user)
		database.DB.First(&user, "user_id = ?", claims["sub"])
		if user.UserID == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// Attach to request
		users := dto.ToDtoUser(user)
		c.Set("user", users)

		// continue
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
