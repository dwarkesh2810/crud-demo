package controller

import (
	categorymodel "crud/internal/modules/category/categoryModel"
	"crud/pkg/database"
	"crud/request"
	"crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategoryList(c *gin.Context) {
	var category []categorymodel.Category
	result := database.DB.Find(&category)
	if result.Error != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "Unexpected database error")
		return
	}
	utils.JsonResponse(c, http.StatusOK, 1, category, "")
}

func CreateCategory(c *gin.Context) {
	var category categorymodel.Category
	var cat request.CategoryRequest

	err := c.ShouldBindJSON(&cat)

	if err != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, utils.ParseError(err)[1])
		return
	}

	category = categorymodel.Category{
		Category_Type: cat.Category,
	}

	result := database.DB.Create(&category)

	if result.Error != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "Unexpected database error or already in List")
		return
	}

	utils.JsonResponse(c, http.StatusOK, 1, category, "")
}

func DeleteCategory(c *gin.Context) {

	var deleteCat request.CategoryDeleteRequest

	err := c.ShouldBindJSON(&deleteCat)

	if err != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, utils.ParseError(err)[1])
		return

	}

	var category categorymodel.Category
	result := database.DB.First(&category, "category_type = ?", deleteCat)

	if result.Error != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "Unexpected database error")
		return
	}

	result = database.DB.Delete(&category)

	if result.Error != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "Unexpected database error")
		return
	}

	utils.JsonResponse(c, http.StatusOK, 1, category, "")
}
