package userroutes

import (
	usercontroller "crud/internal/modules/user/userController"

	"github.com/gin-gonic/gin"
)

// 	router.POST("user/login", controller.Login)

func Routes(router *gin.Engine) {

	userController := usercontroller.New()

	router.POST("user/signup", userController.RegisterUser)
	router.POST("user/login", userController.LoginUser)
	// router.PUT("/post/update", middleware.RequireAuth, controller.UpdatePost)
	// router.DELETE("/post/delete", middleware.RequireAuth, controller.DeletePost)
	// router.GET("/post/by_category", controller.GetPostByCategory)
	// router.GET("/post/by_user", controller.GetPostByUser)
	// router.GET("/post/by_user_cat", controller.GetPostDateByUserAndCategory)
}
