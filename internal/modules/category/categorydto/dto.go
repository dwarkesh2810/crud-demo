package categorydto

import (
	categorymodel "crud/internal/modules/category/categoryModel"
	categoryresponse "crud/internal/modules/category/categoryResponse"
)

func DtOCategoryResponse(cat categorymodel.Category) categoryresponse.CategoryResponse {
	return categoryresponse.CategoryResponse{
		CategoryType : cat.CategoryType,
		ID:       cat.CategoryID,
	}
}

func DtoCategoriesResponse(cat []categorymodel.Category) categoryresponse.CategoriesResponse {
	var response categoryresponse.CategoriesResponse

	for _, c := range cat {
		response.Data = append(response.Data, DtOCategoryResponse(c))
	}

	return response
}
