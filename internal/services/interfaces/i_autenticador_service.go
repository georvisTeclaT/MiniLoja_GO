package interfaces

import (
	"mini-loja/internal/dto"
	"mini-loja/internal/dto/usuario"
)

type IAutenticadorService interface {
	AutenticarUsuario(dados usuario.UsuarioAutenticarDto) dto.ResponseApiDto
}
