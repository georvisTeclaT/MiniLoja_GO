package routers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {

	productRoutes(router, db)
	usuarioRoutes(router, db)

}
