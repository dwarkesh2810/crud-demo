package postroutes

import (
	"crud/internal/middleware"
	postcontroller "crud/internal/modules/post/postController"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	postController := postcontroller.New()

	router.GET("/posts", postController.ShowPost)
	router.POST("/post/create", middleware.RequireAuth, postController.CreatePost)
	router.PUT("/post/update", middleware.RequireAuth, postController.UpdatePost)
	router.DELETE("/post/delete", middleware.RequireAuth, postController.DeletePost)
}
