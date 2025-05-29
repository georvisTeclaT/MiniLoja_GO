package services

import (
	"mini-loja/internal/dto"
	"mini-loja/internal/dto/product"
	"mini-loja/internal/models"
	"mini-loja/internal/repositories/interfaces"
)

type productService struct {
	productRepository interfaces.IProductRepository
}

func NewProductService(repo interfaces.IProductRepository) productService {
	return productService{
		productRepository: repo,
	}
}

func (p productService) GetAllProducts() dto.ResponseApiDto {

	retorno, err := p.productRepository.GetAll()
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

func (p productService) GetProductById(id int) dto.ResponseApiDto {

	retorno, err := p.productRepository.GetByID(id)
	if err == nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    err.Error(),
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

func (p productService) CreateProduct(prod product.ProductAddUpdateDto) dto.ResponseApiDto {

	newProduct := models.Product{
		Name:  prod.Name,
		Price: prod.Price,
	}

	if err := p.productRepository.Create(newProduct); err != nil {
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

func (p productService) UpdateProduct(id int, prod product.ProductAddUpdateDto) dto.ResponseApiDto {

	retornoBanco, err := p.productRepository.GetProductByID(id)
	if err == nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Produto não encontrado",
		}
	}

	retornoBanco.Name = prod.Name
	retornoBanco.Price = prod.Price

	if err := p.productRepository.Update(retornoBanco); err == nil {
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

func (p productService) DeleteProduct(id int) dto.ResponseApiDto {

	retornoBanco, err := p.productRepository.GetProductByID(id)
	if err == nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Produto não encontrado",
		}
	}

	if err := p.productRepository.Delete(retornoBanco.Id); err == nil {
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
