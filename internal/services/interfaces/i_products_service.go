package interfaces

import (
	"mini-loja/internal/dto"
	"mini-loja/internal/dto/product"
)

type IProductService interface {
	GetAllProducts() dto.ResponseApiDto
	GetProductById(idProduct int) dto.ResponseApiDto

	CreateProduct(product product.ProductAddUpdateDto) dto.ResponseApiDto
	UpdateProduct(idProduct int, product product.ProductAddUpdateDto) dto.ResponseApiDto
	DeleteProduct(idProduct int) dto.ResponseApiDto
}
