package routers

import (
	"database/sql"
	"mini-loja/internal/controllers"
	"mini-loja/internal/repositories"
	"mini-loja/internal/services"

	"github.com/gin-gonic/gin"
)

func productRoutes(router *gin.Engine, db *sql.DB) {
	productRepository := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepository)
	productController := controllers.NewProductController(productService)

	productRoutes := router.Group("/product")
	{
		productRoutes.GET("/allProducts", productController.GetAllProducts)
		productRoutes.GET("/productById/:id", productController.GetProductById)

		productRoutes.POST("/productAdd", productController.CreateProduct)
		productRoutes.PUT("/productUpdate/:id", productController.UpdateProduct)
		productRoutes.DELETE("/productDelete/:id", productController.DeleteProduct)
	}
}
