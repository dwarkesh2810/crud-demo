package categoryservice

import (
	categoryrequest "crud/internal/modules/category/categoryRequest"
	categoryresponse "crud/internal/modules/category/categoryResponse"
)

type CategoryServiceInterface interface {
	GetCategories() categoryresponse.CategoriesResponse
	CreateCategory(categoryrequest.CategoryRequest) (categoryresponse.CategoryResponse, error)
}
