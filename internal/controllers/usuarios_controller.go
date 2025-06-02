package controllers

import (
	"mini-loja/internal/dto"
	"mini-loja/internal/dto/usuario"
	"mini-loja/internal/services/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type usuariosController struct {
	_usuarioService interfaces.IUsuarioService
}

func NewUsuarioController(usuarioService interfaces.IUsuarioService) usuariosController {
	return usuariosController{
		_usuarioService: usuarioService,
	}
}

func (u usuariosController) GetAllUsuarios(ctx *gin.Context) {

	products := u._usuarioService.GetAllUsuarios()

	ctx.JSON(http.StatusOK, products)
}

func (u usuariosController) GetUsuarioById(ctx *gin.Context) {

	idParam := ctx.Param("id")

	idUsuario, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	usuario := u._usuarioService.GetUsuarioById(idUsuario)

	ctx.JSON(http.StatusOK, usuario)
}

func (u usuariosController) CreateUsuario(ctx *gin.Context) {

	var input usuario.UsuarioAddUpdateDto

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	newUsuario := usuario.UsuarioAddUpdateDto{
		Nome:      input.Nome,
		Sobrenome: input.Sobrenome,
		Email:     input.Email,
		Telefone:  input.Telefone,
	}

	if input.Nome == "" || input.Sobrenome == "" || input.Email == "" || input.Telefone == "" {
		retorno := dto.ResponseApiDto{
			Status: false,
			Msg:    "Objeto inválido",
		}
		ctx.JSON(http.StatusBadRequest, retorno)
		return
	}

	retornoAddServices := u._usuarioService.CreateUsuario(newUsuario)
	if !retornoAddServices.Status {
		ctx.JSON(http.StatusBadRequest, retornoAddServices)
		return
	}

	ctx.JSON(http.StatusOK, retornoAddServices)
}

func (u usuariosController) UpdateUsuario(ctx *gin.Context) {

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

	if input.Nome == "" || input.Sobrenome == "" || input.Email == "" || input.Telefone == "" {
		retorno := dto.ResponseApiDto{
			Status: false,
			Msg:    "Objeto inválido",
		}
		ctx.JSON(http.StatusBadRequest, retorno)
		return
	}

	updateUsuario := usuario.UsuarioAddUpdateDto{
		Nome:      input.Nome,
		Sobrenome: input.Sobrenome,
		Email:     input.Email,
		Telefone:  input.Telefone,
		Ativo:     input.Ativo,
	}

	retornoUpdateServices := u._usuarioService.UpdateUsuario(idUsuario, updateUsuario)
	if !retornoUpdateServices.Status {
		ctx.JSON(http.StatusBadRequest, retornoUpdateServices)
		return
	}

	ctx.JSON(http.StatusOK, retornoUpdateServices)
}

func (u usuariosController) DeleteUsuario(ctx *gin.Context) {

	idParam := ctx.Param("id")
	idUsuario, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	retornoDeleteServices := u._usuarioService.DeleteUsuario(idUsuario)
	if !retornoDeleteServices.Status {
		ctx.JSON(http.StatusBadRequest, retornoDeleteServices)
		return
	}

	ctx.JSON(http.StatusOK, retornoDeleteServices)
}
