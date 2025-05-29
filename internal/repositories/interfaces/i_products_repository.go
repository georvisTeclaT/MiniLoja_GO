package interfaces

import (
	"mini-loja/internal/dto/product"
	"mini-loja/internal/models"
)

type IProductRepository interface {
	GetAll() ([]product.ProductDto, error)
	GetByID(id int) (product.ProductDto, error)
	GetProductByID(id int) (models.Product, error)

	Create(product models.Product) error
	Update(product models.Product) error
	Delete(id int) error
}
