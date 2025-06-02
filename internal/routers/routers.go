package routers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {

	produtoRoutes(router, db)
	usuarioRoutes(router, db)
	enderecoRoutes(router, db)

}
