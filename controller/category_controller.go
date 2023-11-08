package controller

import (
	"crud/initialise"
	"crud/models"
	"crud/request"
	"crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategoryList(c *gin.Context) {
	var category []models.Category
	result := initialise.DB.Find(&category)
	if result.Error != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "Unexpected database error")
		return
	}
	utils.JsonResponse(c, http.StatusOK, 1, category, "")
}

func CreateCategory(c *gin.Context) {
	var category models.Category
	var cat request.CategoryRequest

	err := c.ShouldBindJSON(&cat)

	if err != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, utils.ParseError(err)[1])
		return
	}

	category = models.Category{
		Category_Type: cat.Category,
	}

	result := initialise.DB.Create(&category)

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

	var category models.Category
	result := initialise.DB.First(&category, "category_type = ?", deleteCat)

	if result.Error != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "Unexpected database error")
		return
	}

	result = initialise.DB.Delete(&category)

	if result.Error != nil {
		utils.JsonResponse(c, http.StatusBadRequest, 0, nil, "Unexpected database error")
		return
	}

	utils.JsonResponse(c, http.StatusOK, 1, category, "")
}
