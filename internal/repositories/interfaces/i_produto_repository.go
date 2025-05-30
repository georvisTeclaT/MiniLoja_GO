package interfaces

import (
	"mini-loja/internal/dto/produto"
	"mini-loja/internal/models"
)

type IProdutoRepository interface {
	GetAll() ([]produto.ProdutoDto, error)
	GetByID(id int) (produto.ProdutoDto, error)
	GetProdutoById(id int) (models.Produto, error)

	Create(produto models.Produto) error
	Update(produto models.Produto) error
	Delete(id int) error
}
