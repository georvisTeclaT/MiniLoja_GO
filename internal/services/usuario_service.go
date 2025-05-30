package services

import (
	"mini-loja/internal/dto"
	"mini-loja/internal/dto/usuario"
	"mini-loja/internal/models"
	"mini-loja/internal/repositories/interfaces"
	"time"
)

type usuarioService struct {
	usuarioRepository interfaces.IUsuarioRepository
}

func NewUsuarioService(repo interfaces.IUsuarioRepository) usuarioService {
	return usuarioService{
		usuarioRepository: repo,
	}
}

func (u usuarioService) GetAllUsuarios() dto.ResponseApiDto {

	retorno, err := u.usuarioRepository.GetAll()
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

	retorno, err := u.usuarioRepository.GetById(id)
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

	newUsuario := models.Usuario{
		Nome:      usuario.Nome,
		Sobrenome: usuario.Sobrenome,
		Email:     usuario.Email,
		Telefone:  usuario.Telefone,
	}

	if err := u.usuarioRepository.Create(newUsuario); err != nil {
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

	retornoBanco, err := u.usuarioRepository.GetUsuarioById(id)
	if err == nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Usuário não encontrado",
		}
	}

	retornoBanco.Nome = usuario.Nome
	retornoBanco.Sobrenome = usuario.Sobrenome
	retornoBanco.Email = usuario.Email
	retornoBanco.Telefone = usuario.Telefone
	retornoBanco.Ativo = usuario.Ativo
	retornoBanco.DataAtualizacao = time.Now()

	if err := u.usuarioRepository.Update(retornoBanco); err == nil {
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

	retornoBanco, err := u.usuarioRepository.GetUsuarioById(id)
	if err == nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Usuário não encontrado",
		}
	}

	if err := u.usuarioRepository.Delete(retornoBanco.Id); err == nil {
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
