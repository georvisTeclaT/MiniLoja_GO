package interfaces

import (
	enderecousuario "mini-loja/internal/dto/endereco_usuario"
	"mini-loja/internal/models"
)

type IEnderecoRepository interface {
	GetAll() ([]enderecousuario.EnderecoUsuarioDto, error)
	GetById(id int) (enderecousuario.EnderecoUsuarioDto, error)
	GetEnderecoById(idEndereco int, idusuario int) (models.EnderecoUsuario, error)

	Create(endereco models.EnderecoUsuario) error
	Update(endereco models.EnderecoUsuario) error
	Delete(id int) error
}
