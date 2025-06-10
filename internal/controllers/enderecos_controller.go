package controllers

import (
	"mini-loja/internal/dto"
	enderecousuario "mini-loja/internal/dto/endereco_usuario"
	"mini-loja/internal/services/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type enderecosController struct {
	_enderecoService interfaces.IEnderecoService
}

func NewEnderecoController(enderecoService interfaces.IEnderecoService) enderecosController {
	return enderecosController{
		_enderecoService: enderecoService,
	}
}

func (e enderecosController) GetAllEnderecos(ctx *gin.Context) {

	produtos := e._enderecoService.GetAllEnderecos()

	ctx.JSON(http.StatusOK, produtos)
}

func (e enderecosController) GetEnderecoById(ctx *gin.Context) {

	idParam := ctx.Param("id")

	idEndereco, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	usuario := e._enderecoService.GetEnderecoById(idEndereco)

	ctx.JSON(http.StatusOK, usuario)
}

func (e enderecosController) CreateEndereco(ctx *gin.Context) {

	idParamUsuario := ctx.Param("idUsuario")
	idUsuario, err := strconv.Atoi(idParamUsuario)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id do usuário inválido"})
		return
	}

	var input enderecousuario.EnderecoUsuarioAddUpdateDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	newEndereco := enderecousuario.EnderecoUsuarioAddUpdateDto{
		NomeRua:     input.NomeRua,
		Numero:      input.Numero,
		Complemento: input.Complemento,
		Bairro:      input.Bairro,
		Cidade:      input.Cidade,
		Estado:      input.Estado,
		Cep:         input.Cep,
	}

	if input.NomeRua == "" || input.Numero == "" || input.Complemento == "" || input.Bairro == "" || input.Cidade == "" || input.Estado == "" || input.Cep == "" {
		retorno := dto.ResponseApiDto{
			Status: false,
			Msg:    "Objeto inválido",
		}
		ctx.JSON(http.StatusBadRequest, retorno)
		return
	}

	retornoAddServices := e._enderecoService.CreateEndereco(idUsuario, newEndereco)
	if !retornoAddServices.Status {
		ctx.JSON(http.StatusBadRequest, retornoAddServices)
		return
	}

	ctx.JSON(http.StatusOK, retornoAddServices)
}

func (e enderecosController) UpdateEndereco(ctx *gin.Context) {

	idParamEndereco := ctx.Param("idEndereco")
	idEndereco, err := strconv.Atoi(idParamEndereco)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id do endereço inválido"})
		return
	}
	idParamUsuario := ctx.Param("idUsuario")
	idUsuario, err := strconv.Atoi(idParamUsuario)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id do usuário inválido"})
		return
	}

	var input enderecousuario.EnderecoUsuarioAddUpdateDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	if input.NomeRua == "" || input.Numero == "" || input.Complemento == "" || input.Bairro == "" || input.Cidade == "" || input.Estado == "" || input.Cep == "" {
		retorno := dto.ResponseApiDto{
			Status: false,
			Msg:    "Objeto inválido",
		}
		ctx.JSON(http.StatusBadRequest, retorno)
		return
	}

	updateEndereco := enderecousuario.EnderecoUsuarioAddUpdateDto{
		NomeRua:     input.NomeRua,
		Numero:      input.Numero,
		Complemento: input.Complemento,
		Bairro:      input.Bairro,
		Cidade:      input.Cidade,
		Estado:      input.Estado,
		Cep:         input.Cep,
	}

	retornoUpdateServices := e._enderecoService.UpdateEndereco(idEndereco, idUsuario, updateEndereco)
	if !retornoUpdateServices.Status {
		ctx.JSON(http.StatusBadRequest, retornoUpdateServices)
		return
	}

	ctx.JSON(http.StatusOK, retornoUpdateServices)
}

func (e enderecosController) DeleteEndereco(ctx *gin.Context) {

	idParamEndereco := ctx.Param("idEndereco")
	idEndereco, err := strconv.Atoi(idParamEndereco)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id do endereço inválido"})
		return
	}
	idParamUsuario := ctx.Param("idUsuario")
	idUsuario, err := strconv.Atoi(idParamUsuario)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id do usuário inválido"})
		return
	}

	retornoDeleteServices := e._enderecoService.DeleteEndereco(idEndereco, idUsuario)
	if !retornoDeleteServices.Status {
		ctx.JSON(http.StatusBadRequest, retornoDeleteServices)
		return
	}

	ctx.JSON(http.StatusOK, retornoDeleteServices)
}
