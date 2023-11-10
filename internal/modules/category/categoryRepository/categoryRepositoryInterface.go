package categoryrepository

import categorymodel "crud/internal/modules/category/categoryModel"

type CategoryRepositoryInterface interface {
	List() []categorymodel.Category
	FindByID(int) (categorymodel.Category, error)
	Create(categorymodel.Category) (categorymodel.Category, error)
	FindByCategory(string) (categorymodel.Category, error)
}
