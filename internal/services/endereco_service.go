package services

import (
	"mini-loja/internal/dto"
	enderecousuario "mini-loja/internal/dto/endereco_usuario"
	"mini-loja/internal/models"
	"mini-loja/internal/repositories/interfaces"
	"time"
)

type enderecoService struct {
	_enderecoRepository interfaces.IEnderecoRepository
	_usuarioRepository  interfaces.IUsuarioRepository
}

func NewEnderecoService(enderecoRepository interfaces.IEnderecoRepository, usuarioRepository interfaces.IUsuarioRepository) enderecoService {
	return enderecoService{
		_enderecoRepository: enderecoRepository,
		_usuarioRepository:  usuarioRepository,
	}
}

func (e enderecoService) GetAllEnderecos() dto.ResponseApiDto {
	retorno, err := e._enderecoRepository.GetAll()
	if err != nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    err.Error(),
		}
	} else if len(retorno) <= 0 {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Não existem registros de endereço no banco de dados",
		}
	}

	return dto.ResponseApiDto{
		Status: true,
		Msg:    "Dados do endereço retornados com sucesso",
		Data:   retorno,
	}
}

func (e enderecoService) GetEnderecoById(id int) dto.ResponseApiDto {

	retorno, err := e._enderecoRepository.GetById(id)
	if err == nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Erro de sistema",
		}
	} else if retorno.Id == 0 {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Não existem registros de endereço no banco de dados",
		}
	}

	return dto.ResponseApiDto{
		Status: true,
		Msg:    "Dados de endereço retornados com sucesso",
		Data:   retorno,
	}
}

func (e enderecoService) CreateEndereco(idUsuario int, endereco enderecousuario.EnderecoUsuarioAddUpdateDto) dto.ResponseApiDto {

	retornoUsuarioBanco, err := e._usuarioRepository.GetUsuarioById(idUsuario)
	if err == nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Usuário não encontrado",
		}
	}

	newEndereco := models.EnderecoUsuario{
		UsuarioId:   retornoUsuarioBanco.Id,
		NomeRua:     endereco.NomeRua,
		Numero:      endereco.Numero,
		Complemento: endereco.Complemento,
		Bairro:      endereco.Bairro,
		Cidade:      endereco.Cidade,
		Estado:      endereco.Estado,
		Cep:         endereco.Cep,
	}

	if err := e._enderecoRepository.Create(newEndereco); err != nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Erro ao inserir os dados do endereço no banco de dados",
		}
	}

	return dto.ResponseApiDto{
		Status: true,
		Msg:    "Endereço inserido com sucesso",
	}
}

func (e enderecoService) UpdateEndereco(idEndereco int, idUsuario int, endereco enderecousuario.EnderecoUsuarioAddUpdateDto) dto.ResponseApiDto {

	retornoBanco, err := e._enderecoRepository.GetEnderecoById(idEndereco, idUsuario)
	if err == nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Endereço não encontrado",
		}
	}

	retornoBanco.NomeRua = endereco.NomeRua
	retornoBanco.Numero = endereco.Numero
	retornoBanco.Complemento = endereco.Complemento
	retornoBanco.Bairro = endereco.Bairro
	retornoBanco.Cidade = endereco.Cidade
	retornoBanco.Estado = endereco.Estado
	retornoBanco.Cep = endereco.Cep
	retornoBanco.DataAtualizacao = time.Now()

	if err := e._enderecoRepository.Update(retornoBanco); err == nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Erro ao atualizar os dados do endereço no banco de dados",
		}
	}

	return dto.ResponseApiDto{
		Status: true,
		Msg:    "Endereço atualizado com sucesso",
	}
}

func (e enderecoService) DeleteEndereco(idEndereco int, idUsuario int) dto.ResponseApiDto {

	retornoBanco, err := e._enderecoRepository.GetEnderecoById(idEndereco, idUsuario)
	if err == nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Endereço não encontrado",
		}
	}

	if err := e._enderecoRepository.Delete(retornoBanco.Id); err == nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Erro ao deletar os dados do endereço no banco de dados",
		}
	}

	return dto.ResponseApiDto{
		Status: true,
		Msg:    "Endereço deletado com sucesso",
	}
}
