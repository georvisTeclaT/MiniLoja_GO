package controllers

import (
	"mini-loja/internal/dto"
	"mini-loja/internal/dto/product"
	"mini-loja/internal/services/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductsController struct {
	productService interfaces.IProductService
}

func NewProductController(s interfaces.IProductService) *ProductsController {
	return &ProductsController{s}
}

func (p ProductsController) GetAllProducts(ctx *gin.Context) {

	products := p.productService.GetAllProducts()

	ctx.JSON(http.StatusOK, products)
}

func (c ProductsController) GetProductById(ctx *gin.Context) {

	idParam := ctx.Param("id")

	idProduct, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	product := c.productService.GetProductById(idProduct)

	ctx.JSON(http.StatusOK, product)
}

func (p ProductsController) CreateProduct(ctx *gin.Context) {

	var input product.ProductAddUpdateDto

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	newProduct := product.ProductAddUpdateDto{
		Name:  input.Name,
		Price: input.Price,
	}

	if input.Name == "" || input.Price <= 0 {
		retorno := dto.ResponseApiDto{
			Status: false,
			Msg:    "Os campos Name e Price são obrigatórios",
		}
		ctx.JSON(http.StatusBadRequest, retorno)
		return
	}

	retornoAddServices := p.productService.CreateProduct(newProduct)
	if !retornoAddServices.Status {
		ctx.JSON(http.StatusBadRequest, retornoAddServices)
		return
	}

	ctx.JSON(http.StatusOK, retornoAddServices)
}

func (p ProductsController) UpdateProduct(ctx *gin.Context) {

	idParam := ctx.Param("id")
	idProduct, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input product.ProductAddUpdateDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	if input.Name == "" || input.Price <= 0 {
		retorno := dto.ResponseApiDto{
			Status: false,
			Msg:    "Os campos Name e Price são obrigatórios",
		}
		ctx.JSON(http.StatusBadRequest, retorno)
		return
	}

	updateProduct := product.ProductAddUpdateDto{
		Name:  input.Name,
		Price: input.Price,
	}

	retornoUpdateServices := p.productService.UpdateProduct(idProduct, updateProduct)
	if !retornoUpdateServices.Status {
		ctx.JSON(http.StatusBadRequest, retornoUpdateServices)
		return
	}

	ctx.JSON(http.StatusOK, retornoUpdateServices)
}

func (p ProductsController) DeleteProduct(ctx *gin.Context) {

	idParam := ctx.Param("id")
	idProduct, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	retornoDeleteServices := p.productService.DeleteProduct(idProduct)
	if !retornoDeleteServices.Status {
		ctx.JSON(http.StatusBadRequest, retornoDeleteServices)
		return
	}

	ctx.JSON(http.StatusOK, retornoDeleteServices)
}
