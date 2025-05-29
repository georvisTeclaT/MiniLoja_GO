package interfaces

import (
	"mini-loja/internal/dto/usuario"
	"mini-loja/internal/models"
)

type IUsuarioRepository interface {
	GetAll() ([]usuario.UsuarioDto, error)
	GetByID(id int) (usuario.UsuarioDto, error)
	GetUsuarioByID(id int) (models.Usuario, error)

	Create(product models.Usuario) error
	Update(product models.Usuario) error
	Delete(id int) error
}
