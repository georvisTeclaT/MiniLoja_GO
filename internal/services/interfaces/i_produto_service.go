package interfaces

import (
	"mini-loja/internal/dto"
	"mini-loja/internal/dto/produto"
)

type IProdutoService interface {
	GetAllProdutos() dto.ResponseApiDto
	GetProdutoById(idProduto int) dto.ResponseApiDto

	CreateProduto(produto produto.ProdutoAddUpdateDto) dto.ResponseApiDto
	UpdateProduto(idProduto int, produto produto.ProdutoAddUpdateDto) dto.ResponseApiDto
	DeleteProduto(idProduto int) dto.ResponseApiDto
}
