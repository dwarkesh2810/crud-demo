package categoryroutes

import (
	categorycontroller "crud/internal/modules/category/categoryController"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	categoryController:= categorycontroller.New()
	router.GET("/category", categoryController.GetCategories)
	router.POST("/category/create", categoryController.CreateCategory)
	// router.DELETE("/category/delete", controller.DeleteCategory)

}
