package routers

import (
	"database/sql"
	"mini-loja/internal/controllers"
	"mini-loja/internal/repositories"
	"mini-loja/internal/services"

	"github.com/gin-gonic/gin"
)

func autenticadorRoutes(router *gin.Engine, db *sql.DB) {

	// Repositories
	usuarioRepository := repositories.NewUsuarioRepository(db)

	// Services
	autenticadorService := services.NewAutenticadorService(usuarioRepository)

	// Controller
	autenticadorController := controllers.NewAutenticadorController(autenticadorService)

	autenticadorRoutes := router.Group("/autenticador")
	{
		autenticadorRoutes.POST("/autenticarUsuario", autenticadorController.AutenticarUsuario)
	}
}
