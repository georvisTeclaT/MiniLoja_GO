package routers

import (
	"database/sql"
	"mini-loja/internal/controllers"
	"mini-loja/internal/repositories"
	"mini-loja/internal/services"

	"github.com/gin-gonic/gin"
)

func usuarioRoutes(router *gin.Engine, db *sql.DB) {

	// Repositories
	usuarioRepository := repositories.NewUsuarioRepository(db)

	// Services
	usuarioService := services.NewUsuarioService(usuarioRepository)

	// Controller
	usuariosController := controllers.NewUsuarioController(usuarioService)

	usuarioRoutes := router.Group("/usuario")
	{
		usuarioRoutes.GET("/allUsuarios", usuariosController.GetAllUsuarios)
		usuarioRoutes.GET("/usuarioById/:id", usuariosController.GetUsuarioById)

		usuarioRoutes.POST("/usuarioAdd", usuariosController.CreateUsuario)
		usuarioRoutes.PUT("/usuarioUpdate/:id", usuariosController.UpdateUsuario)
		usuarioRoutes.DELETE("/usuarioDelete/:id", usuariosController.DeleteUsuario)
	}
}
