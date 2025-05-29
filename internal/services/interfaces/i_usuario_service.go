package interfaces

import (
	"mini-loja/internal/dto"
	"mini-loja/internal/dto/usuario"
)

type IUsuarioService interface {
	GetAllUsuarios() dto.ResponseApiDto
	GetUsuarioById(idUsuario int) dto.ResponseApiDto

	CreateUsuario(usuario usuario.UsuarioAddUpdateDto) dto.ResponseApiDto
	UpdateUsuario(idUsuario int, usuario usuario.UsuarioAddUpdateDto) dto.ResponseApiDto
	DeleteUsuario(idUsuario int) dto.ResponseApiDto
}
