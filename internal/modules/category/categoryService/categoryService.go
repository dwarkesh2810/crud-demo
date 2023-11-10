package categoryservice

import (
	categorymodel "crud/internal/modules/category/categoryModel"
	categoryrepository "crud/internal/modules/category/categoryRepository"
	categoryrequest "crud/internal/modules/category/categoryRequest"
	categoryresponse "crud/internal/modules/category/categoryResponse"
	"crud/internal/modules/category/categorydto"
	"errors"
)

type CategoryService struct {
	categoryRepo categoryrepository.CategoryRepositoryInterface
}

func New() *CategoryService {
	return &CategoryService{
		categoryRepo: categoryrepository.New(),
	}
}

func (categoryService *CategoryService) CreateCategory(request categoryrequest.CategoryRequest) (categoryresponse.CategoryResponse, error) {
	var category categorymodel.Category

	cat, _ := categoryService.categoryRepo.FindByCategory(request.Category)

	if cat.CategoryType != "" {
		return categoryresponse.CategoryResponse{}, errors.New("already added this category")
	}

	category.CategoryType = request.Category

	cat, err := categoryService.categoryRepo.Create(category)

	if err != nil {
		return categoryresponse.CategoryResponse{}, err
	}

	return categorydto.DtOCategoryResponse(cat), nil
}

func (categoryService *CategoryService) GetCategories() categoryresponse.CategoriesResponse {
	posts := categoryService.categoryRepo.List()
	return categorydto.DtoCategoriesResponse(posts)
}
