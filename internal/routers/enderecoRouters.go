package routers

import (
	"database/sql"
	"mini-loja/internal/controllers"
	"mini-loja/internal/repositories"
	"mini-loja/internal/services"

	"github.com/gin-gonic/gin"
)

func enderecoRoutes(router *gin.Engine, db *sql.DB) {

	// Repositories
	enderecoRepository := repositories.NewEnderecoRepository(db)
	usuarioRepository := repositories.NewUsuarioRepository(db)

	// Services
	enderecoService := services.NewEnderecoService(enderecoRepository, usuarioRepository)

	// Controller
	enderecoController := controllers.NewEnderecoController(enderecoService)

	enderecoRoutes := router.Group("/endereco")
	{
		enderecoRoutes.GET("/allEnderecos", enderecoController.GetAllEnderecos)
		enderecoRoutes.GET("/enderecoById/:id", enderecoController.GetEnderecoById)
		enderecoRoutes.POST("/enderecoAdd/:idUsuario", enderecoController.CreateEndereco)
		enderecoRoutes.PUT("/enderecoUpdate/:idEndereco/:idUsuario", enderecoController.UpdateEndereco)
		enderecoRoutes.DELETE("/enderecoDelete/:idEndereco/:idUsuario", enderecoController.DeleteEndereco)
	}
}
