package services

import (
	"mini-loja/internal/dto"
	"mini-loja/internal/dto/usuario"
	"mini-loja/internal/models"
	"mini-loja/internal/repositories/interfaces"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type usuarioService struct {
	_usuarioRepository interfaces.IUsuarioRepository
}

func NewUsuarioService(usuarioRepository interfaces.IUsuarioRepository) usuarioService {
	return usuarioService{
		_usuarioRepository: usuarioRepository,
	}
}

func (u usuarioService) GetAllUsuarios() dto.ResponseApiDto {

	retorno, err := u._usuarioRepository.GetAll()
	if err != nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    err.Error(),
		}
	} else if len(retorno) <= 0 {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Não existem registros de usuários no banco de dados",
		}
	}

	return dto.ResponseApiDto{
		Status: true,
		Msg:    "Dados do usuário retornados com sucesso",
		Data:   retorno,
	}
}

func (u usuarioService) GetUsuarioById(id int) dto.ResponseApiDto {

	retorno, err := u._usuarioRepository.GetById(id)
	if err == nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Erro de sistema",
		}
	} else if retorno.Id == 0 {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Não existem registros de usuários no banco de dados",
		}
	}

	return dto.ResponseApiDto{
		Status: true,
		Msg:    "Dados de usuários retornados com sucesso",
		Data:   retorno,
	}
}

func (u usuarioService) CreateUsuario(usuario usuario.UsuarioAddUpdateDto) dto.ResponseApiDto {

	_senha, err := HashPassword(usuario.Senha)
	if err != nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    err.Error(),
		}
	}

	newUsuario := models.Usuario{
		Nome:      usuario.Nome,
		Sobrenome: usuario.Sobrenome,
		Email:     usuario.Email,
		Telefone:  usuario.Telefone,
		Senha:     _senha,
	}

	if err := u._usuarioRepository.Create(newUsuario); err != nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Erro ao inserir os dados do usuário no banco de dados",
		}
	}

	return dto.ResponseApiDto{
		Status: true,
		Msg:    "Usuário inserido com sucesso",
	}
}

func (u usuarioService) UpdateUsuario(id int, usuario usuario.UsuarioAddUpdateDto) dto.ResponseApiDto {

	retornoBanco, err := u._usuarioRepository.GetUsuarioById(id)
	if err != nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Usuário não encontrado",
		}
	}

	_senha, err := HashPassword(usuario.Senha)
	if err != nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    err.Error(),
		}
	}

	retornoBanco.Nome = usuario.Nome
	retornoBanco.Sobrenome = usuario.Sobrenome
	retornoBanco.Email = usuario.Email
	retornoBanco.Telefone = usuario.Telefone
	retornoBanco.Ativo = usuario.Ativo
	retornoBanco.DataAtualizacao = time.Now()
	retornoBanco.Senha = _senha

	if err := u._usuarioRepository.Update(retornoBanco); err == nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Erro ao atualizar os dados do usuário no banco de dados",
		}
	}

	return dto.ResponseApiDto{
		Status: true,
		Msg:    "Usuário atualizado com sucesso",
	}
}

func (u usuarioService) DeleteUsuario(id int) dto.ResponseApiDto {

	retornoBanco, err := u._usuarioRepository.GetUsuarioById(id)
	if err == nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Usuário não encontrado",
		}
	}

	if err := u._usuarioRepository.Delete(retornoBanco.Id); err == nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Erro ao deletar os dados do usuário no banco de dados",
		}
	}

	return dto.ResponseApiDto{
		Status: true,
		Msg:    "Usuário deletado com sucesso",
	}
}

// Uteis

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
