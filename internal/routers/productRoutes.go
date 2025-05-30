package routers

import (
	"database/sql"
	"mini-loja/internal/controllers"
	"mini-loja/internal/repositories"
	"mini-loja/internal/services"

	"github.com/gin-gonic/gin"
)

func produtoRoutes(router *gin.Engine, db *sql.DB) {

	produtoRepository := repositories.NewProdutoRepository(db)
	produtoService := services.NewProdutoService(produtoRepository)
	produtoController := controllers.NewProdutoController(produtoService)

	produtoRoutes := router.Group("/produto")
	{
		produtoRoutes.GET("/allProdutos", produtoController.GetAllProdutos)
		produtoRoutes.GET("/produtoById/:id", produtoController.GetProdutoById)

		produtoRoutes.POST("/produtoAdd", produtoController.CreateProduto)
		produtoRoutes.PUT("/produtoUpdate/:id", produtoController.UpdateProduto)
		produtoRoutes.DELETE("/produtoDelete/:id", produtoController.DeleteProduto)
	}
}
