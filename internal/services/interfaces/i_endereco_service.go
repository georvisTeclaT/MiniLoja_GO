package interfaces

import (
	"mini-loja/internal/dto"
	enderecousuario "mini-loja/internal/dto/endereco_usuario"
)

type IEnderecoService interface {
	GetAllEnderecos() dto.ResponseApiDto
	GetEnderecoById(idEndereco int) dto.ResponseApiDto

	CreateEndereco(idUsuario int, endereco enderecousuario.EnderecoUsuarioAddUpdateDto) dto.ResponseApiDto
	UpdateEndereco(idEndereco int, idUsuario int, endereco enderecousuario.EnderecoUsuarioAddUpdateDto) dto.ResponseApiDto
	DeleteEndereco(idEndereco int, idUsuario int) dto.ResponseApiDto
}
