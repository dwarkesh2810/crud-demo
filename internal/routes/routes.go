package routes

import (
	postroutes "crud/internal/modules/post/postRoutes"
	userroutes "crud/internal/modules/user/userRoutes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	postroutes.Routes(router)
	userroutes.Routes(router)
}
