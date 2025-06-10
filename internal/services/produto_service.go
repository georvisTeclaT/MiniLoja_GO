package services

import (
	"mini-loja/internal/dto"
	"mini-loja/internal/dto/produto"
	"mini-loja/internal/models"
	"mini-loja/internal/repositories/interfaces"
	"time"
)

type produtoService struct {
	_produtoRepository interfaces.IProdutoRepository
}

func NewProdutoService(produtoRepository interfaces.IProdutoRepository) produtoService {
	return produtoService{
		_produtoRepository: produtoRepository,
	}
}

func (p produtoService) GetAllProdutos() dto.ResponseApiDto {

	retorno, err := p._produtoRepository.GetAll()
	if err != nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    err.Error(),
		}
	} else if len(retorno) <= 0 {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Não existem registros de produtos no banco de dados",
		}
	}

	return dto.ResponseApiDto{
		Status: true,
		Msg:    "Dados dos produtos retornados com sucesso",
		Data:   retorno,
	}

}

func (p produtoService) GetProdutoById(id int) dto.ResponseApiDto {

	retorno, err := p._produtoRepository.GetByID(id)
	if err == nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Erro de sistema",
		}
	} else if retorno.Id == 0 {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Produto não encontrado",
		}
	}

	return dto.ResponseApiDto{
		Status: true,
		Msg:    "Dados do produto retornados com sucesso",
		Data:   retorno,
	}

}

func (p produtoService) CreateProduto(prod produto.ProdutoAddUpdateDto) dto.ResponseApiDto {

	newProduto := models.Produto{
		Nome:      prod.Nome,
		Descricao: prod.Descricao,
		Stock:     prod.Stock,
		Preco:     prod.Preco,
	}

	if err := p._produtoRepository.Create(newProduto); err != nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Erro ao inserir o produto no banco de dados",
		}
	}

	return dto.ResponseApiDto{
		Status: true,
		Msg:    "Produto inserido com sucesso",
	}
}

func (p produtoService) UpdateProduto(id int, prod produto.ProdutoAddUpdateDto) dto.ResponseApiDto {

	retornoBanco, err := p._produtoRepository.GetProdutoById(id)
	if err == nil || retornoBanco.Id <= 0 {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Produto não encontrado",
		}
	}

	retornoBanco.Nome = prod.Nome
	retornoBanco.Descricao = prod.Descricao
	retornoBanco.Stock = prod.Stock
	retornoBanco.Preco = prod.Preco
	retornoBanco.DataAtualizacao = time.Now()

	if err := p._produtoRepository.Update(retornoBanco); err == nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Erro ao atualizar o produto no banco de dados",
		}
	}

	return dto.ResponseApiDto{
		Status: true,
		Msg:    "Produto atualizado com sucesso",
	}
}

func (p produtoService) DeleteProduto(id int) dto.ResponseApiDto {

	retornoBanco, err := p._produtoRepository.GetProdutoById(id)
	if err == nil || retornoBanco.Id <= 0 {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Produto não encontrado",
		}
	}

	if err := p._produtoRepository.Delete(retornoBanco.Id); err == nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Erro ao deletar o produto no banco de dados",
		}
	}

	return dto.ResponseApiDto{
		Status: true,
		Msg:    "Produto deletado com sucesso",
	}
}
