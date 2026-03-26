package repository

import "github.com/wwwmonster/eShopApp/go/v2/internal/domain"

type CatalogRepository interface {
	CreateCategory(c *domain.Category) error
	FindCategories() ([]*domain.Category, error)
	FindCategoryById(id int) (*domain.Category, error)
	EditCategory(c *domain.Category) (*domain.Category, error)
	DeleteCategory(id int) error

	CreateProduct(*domain.Product) error
	FindProducts() ([]*domain.Product, error)
	FindProductById(id int) (*domain.Product, error)
	FindSellerProducts(id int) ([]*domain.Product, error)
	EditProduct(e *domain.Product) (*domain.Product, error)
	DeleteProduct(e *domain.Product) error
}
