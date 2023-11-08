package categoryservice

import (
	categoryrepository "crud/internal/modules/category/categoryRepository"
)

type CategoryService struct {
	categoryRepo categoryrepository.CategoryRepositoryInterface
}

func New() *CategoryService {
	return &CategoryService{
		categoryRepo: categoryrepository.New(),
	}
}
