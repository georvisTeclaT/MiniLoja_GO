package interfaces

import (
	"mini-loja/internal/dto/usuario"
	"mini-loja/internal/models"
)

type IUsuarioRepository interface {
	GetAll() ([]usuario.UsuarioDto, error)
	GetById(id int) (usuario.UsuarioDto, error)
	GetUsuarioById(id int) (models.Usuario, error)

	Create(usuario models.Usuario) error
	Update(usuario models.Usuario) error
	Delete(id int) error
}
