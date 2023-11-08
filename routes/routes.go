package routes

import (
	"crud/controller"
	"crud/middleware"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Init() {
	router = gin.Default()
}

func GetRouter() *gin.Engine {
	return router
}

func Routes(router *gin.Engine) {
	router.GET("/posts", controller.GetPosts)
	router.POST("/post/create", middleware.RequireAuth, controller.CreatePost)
	router.PUT("/post/update", middleware.RequireAuth, controller.UpdatePost)
	router.DELETE("/post/delete", middleware.RequireAuth, controller.DeletePost)

	router.POST("user/signup", controller.Sugnup)
	router.POST("user/login", controller.Login)
	router.GET("/category", controller.GetCategoryList)
	router.POST("/category/create", controller.CreateCategory)
	router.DELETE("/category/delete", controller.DeleteCategory)
}

func RegisterRoutes() {

	Routes(GetRouter())
}

func Serve() {
	r := GetRouter()

	err := r.Run(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))

	if err != nil {
		log.Fatal("error in routing")
		return
	}
}
