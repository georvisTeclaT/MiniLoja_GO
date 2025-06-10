package controllers

import (
	"mini-loja/internal/dto"
	"mini-loja/internal/dto/usuario"
	"mini-loja/internal/services/interfaces"
	"net"
	"net/http"
	"net/mail"
	"strings"

	"github.com/gin-gonic/gin"
)

type autenticadorController struct {
	_autenticadorService interfaces.IAutenticadorService
}

func NewAutenticadorController(autenticadorService interfaces.IAutenticadorService) autenticadorController {
	return autenticadorController{
		_autenticadorService: autenticadorService,
	}
}

func (a autenticadorController) AutenticarUsuario(ctx *gin.Context) {

	var input usuario.UsuarioAutenticarDto

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}
	retornoErro := dto.ResponseApiDto{
		Status: false,
		Msg:    "",
	}

	isValidEmail := isValidEmail(input.Email)
	if input.Email == "" || !isValidEmail {
		retornoErro.Msg = "E-mail inválido"
		ctx.JSON(http.StatusBadRequest, retornoErro)
		return
	}

	if input.Senha == "" || len(input.Senha) < 4 {
		retornoErro.Msg = "Senha inválida"
		ctx.JSON(http.StatusBadRequest, retornoErro)
		return
	}

	retornoAutenticador := a._autenticadorService.AutenticarUsuario(input)
	if !retornoAutenticador.Status {
		ctx.JSON(http.StatusBadRequest, retornoAutenticador)
		return
	}

	ctx.JSON(http.StatusOK, retornoAutenticador)
}

func isValidEmail(addr string) bool {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return false
	}

	domain := a.Address[strings.LastIndex(a.Address, "@")+1:]
	mx, err := net.LookupMX(domain)
	return err == nil && len(mx) > 0
}
