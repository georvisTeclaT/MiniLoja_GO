package controllers

import (
	"mini-loja/internal/dto"
	"mini-loja/internal/dto/usuario"
	"mini-loja/internal/services/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UsuariosController struct {
	usuarioService interfaces.IUsuarioService
}

func NewUsuarioController(s interfaces.IUsuarioService) *UsuariosController {
	return &UsuariosController{s}
}

func (u UsuariosController) GetAllUsuarios(ctx *gin.Context) {

	products := u.usuarioService.GetAllUsuarios()

	ctx.JSON(http.StatusOK, products)
}

func (u UsuariosController) GetUsuarioById(ctx *gin.Context) {

	idParam := ctx.Param("id")

	idUsuario, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	usuario := u.usuarioService.GetUsuarioById(idUsuario)

	ctx.JSON(http.StatusOK, usuario)
}

func (u UsuariosController) CreateUsuario(ctx *gin.Context) {

	var input usuario.UsuarioAddUpdateDto

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	newUsuario := usuario.UsuarioAddUpdateDto{
		Nome:      input.Nome,
		Sobrenome: input.Sobrenome,
	}

	if input.Nome == "" || input.Sobrenome == "" {
		retorno := dto.ResponseApiDto{
			Status: false,
			Msg:    "Os campos Nome e Sobrenome são obrigatórios",
		}
		ctx.JSON(http.StatusBadRequest, retorno)
		return
	}

	retornoAddServices := u.usuarioService.CreateUsuario(newUsuario)
	if !retornoAddServices.Status {
		ctx.JSON(http.StatusBadRequest, retornoAddServices)
		return
	}

	ctx.JSON(http.StatusOK, retornoAddServices)
}

func (u UsuariosController) UpdateUsuario(ctx *gin.Context) {

	idParam := ctx.Param("id")
	idUsuario, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input usuario.UsuarioAddUpdateDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	if input.Nome == "" || input.Sobrenome == "" {
		retorno := dto.ResponseApiDto{
			Status: false,
			Msg:    "Os campos Nome e Sobrenome são obrigatórios",
		}
		ctx.JSON(http.StatusBadRequest, retorno)
		return
	}

	updateUsuario := usuario.UsuarioAddUpdateDto{
		Nome:      input.Nome,
		Sobrenome: input.Sobrenome,
	}

	retornoUpdateServices := u.usuarioService.UpdateUsuario(idUsuario, updateUsuario)
	if !retornoUpdateServices.Status {
		ctx.JSON(http.StatusBadRequest, retornoUpdateServices)
		return
	}

	ctx.JSON(http.StatusOK, retornoUpdateServices)
}

func (u UsuariosController) DeleteUsuario(ctx *gin.Context) {

	idParam := ctx.Param("id")
	idUsuario, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	retornoDeleteServices := u.usuarioService.DeleteUsuario(idUsuario)
	if !retornoDeleteServices.Status {
		ctx.JSON(http.StatusBadRequest, retornoDeleteServices)
		return
	}

	ctx.JSON(http.StatusOK, retornoDeleteServices)
}
