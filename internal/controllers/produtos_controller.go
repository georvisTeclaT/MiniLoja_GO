package controllers

import (
	"mini-loja/internal/dto"
	"mini-loja/internal/dto/produto"
	"mini-loja/internal/services/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProdutosController struct {
	produtoService interfaces.IProdutoService
}

func NewProdutoController(s interfaces.IProdutoService) *ProdutosController {
	return &ProdutosController{s}
}

func (p ProdutosController) GetAllProdutos(ctx *gin.Context) {

	produtos := p.produtoService.GetAllProdutos()

	ctx.JSON(http.StatusOK, produtos)
}

func (c ProdutosController) GetProdutoById(ctx *gin.Context) {

	idParam := ctx.Param("id")

	idProduto, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	produto := c.produtoService.GetProdutoById(idProduto)

	ctx.JSON(http.StatusOK, produto)
}

func (p ProdutosController) CreateProduto(ctx *gin.Context) {

	var input produto.ProdutoAddUpdateDto

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	if input.Nome == "" || input.Descricao == "" || input.Stock <= 0 || input.Preco <= 0 {
		retorno := dto.ResponseApiDto{
			Status: false,
			Msg:    "Objeto inválido",
		}
		ctx.JSON(http.StatusBadRequest, retorno)
		return
	}

	newProduto := produto.ProdutoAddUpdateDto{
		Nome:      input.Nome,
		Descricao: input.Descricao,
		Stock:     input.Stock,
		Preco:     input.Preco,
	}

	retornoAddServices := p.produtoService.CreateProduto(newProduto)
	if !retornoAddServices.Status {
		ctx.JSON(http.StatusBadRequest, retornoAddServices)
		return
	}

	ctx.JSON(http.StatusOK, retornoAddServices)
}

func (p ProdutosController) UpdateProduto(ctx *gin.Context) {

	idParam := ctx.Param("id")
	idProduto, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input produto.ProdutoAddUpdateDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	if input.Nome == "" || input.Descricao == "" || input.Stock <= 0 || input.Preco <= 0 {
		retorno := dto.ResponseApiDto{
			Status: false,
			Msg:    "Objeto inválido",
		}
		ctx.JSON(http.StatusBadRequest, retorno)
		return
	}

	updateProduto := produto.ProdutoAddUpdateDto{
		Nome:      input.Nome,
		Descricao: input.Descricao,
		Stock:     input.Stock,
		Preco:     input.Preco,
	}

	retornoUpdateServices := p.produtoService.UpdateProduto(idProduto, updateProduto)
	if !retornoUpdateServices.Status {
		ctx.JSON(http.StatusBadRequest, retornoUpdateServices)
		return
	}

	ctx.JSON(http.StatusOK, retornoUpdateServices)
}

func (p ProdutosController) DeleteProduto(ctx *gin.Context) {

	idParam := ctx.Param("id")
	idProduto, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	retornoDeleteServices := p.produtoService.DeleteProduto(idProduto)
	if !retornoDeleteServices.Status {
		ctx.JSON(http.StatusBadRequest, retornoDeleteServices)
		return
	}

	ctx.JSON(http.StatusOK, retornoDeleteServices)
}
