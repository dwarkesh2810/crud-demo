package categorycontroller

import categoryservice "crud/internal/modules/category/categoryService"

type Controllers struct {
	categoryService categoryservice.CategoryServiceInterface
}

func New() *Controllers {
	return &Controllers{
		categoryService: categoryservice.New(),
	}
}
