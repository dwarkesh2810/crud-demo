package categorycontroller

import (
	categoryrequest "crud/internal/modules/category/categoryRequest"
	categoryservice "crud/internal/modules/category/categoryService"
	"crud/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryControllers struct {
	categoryService categoryservice.CategoryServiceInterface
}

func New() *CategoryControllers {
	return &CategoryControllers{
		categoryService: categoryservice.New(),
	}
}

func (categoryControllers *CategoryControllers) CreateCategory(c *gin.Context) {
	var request categoryrequest.CategoryRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		helper.JsonResponse(c, http.StatusBadRequest, 0, nil, helper.ParseError(err)[1])
		return
	}

	cat, err := categoryControllers.categoryService.CreateCategory(request)

	if err != nil {
		helper.JsonResponse(c, http.StatusBadRequest, 0, nil, err.Error())
		return
	}

	helper.JsonResponse(c, http.StatusOK, 1, cat, "")
}

func (categoryControllers *CategoryControllers) GetCategories(c *gin.Context) {
	categories := categoryControllers.categoryService.GetCategories()
	helper.JsonResponse(c, http.StatusOK, 1, categories, "")
}
