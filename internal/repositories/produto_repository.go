package repositories

import (
	"database/sql"
	"mini-loja/internal/dto/produto"
	"mini-loja/internal/models"
	"mini-loja/internal/repositories/interfaces"
	"time"
)

type produtoRepository struct {
	db *sql.DB
}

func NewProdutoRepository(db *sql.DB) interfaces.IProdutoRepository {
	return produtoRepository{db: db}
}

func (p produtoRepository) GetAll() ([]produto.ProdutoDto, error) {
	rows, err := p.db.Query("SELECT id, nome, descricao, stock, preco, ativo, data_criacao FROM produto")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var produtoLista []produto.ProdutoDto
	for rows.Next() {
		var produto produto.ProdutoDto
		var dataCriacao time.Time
		if err := rows.Scan(
			&produto.Id,
			&produto.Nome,
			&produto.Descricao,
			&produto.Stock,
			&produto.Preco,
			&produto.Ativo,
			&dataCriacao); err != nil {
			return nil, err
		}

		produto.DataCriacao = dataCriacao.Format("02/01/2006")

		produtoLista = append(produtoLista, produto)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return produtoLista, nil
}

func (p produtoRepository) GetByID(id int) (produto.ProdutoDto, error) {
	rows, err := p.db.Query("SELECT id, nome, descricao, stock, preco, ativo, data_criacao FROM produto WHERE id = $1", id)
	if err != nil {
		return produto.ProdutoDto{}, err
	}
	defer rows.Close()

	var produto produto.ProdutoDto
	if rows.Next() {
		var dataCriacao time.Time
		if err := rows.Scan(
			&produto.Id,
			&produto.Nome,
			&produto.Descricao,
			&produto.Stock,
			&produto.Preco,
			&produto.Ativo,
			&dataCriacao); err != nil {
			return produto, err
		}

		produto.DataCriacao = dataCriacao.Format("02/01/2006")
	}

	return produto, sql.ErrNoRows
}

func (p produtoRepository) GetProdutoById(id int) (models.Produto, error) {
	rows, err := p.db.Query("SELECT id, nome, descricao, stock, preco, ativo, data_criacao, data_atualizacao FROM produto WHERE id = $1", id)
	if err != nil {
		return models.Produto{}, err
	}
	defer rows.Close()

	var produto models.Produto
	if rows.Next() {
		if err := rows.Scan(
			&produto.Id,
			&produto.Nome,
			&produto.Descricao,
			&produto.Stock,
			&produto.Preco,
			&produto.Ativo,
			&produto.DataCriacao,
			&produto.DataAtualizacao); err != nil {
			return produto, err
		}
	}

	return produto, sql.ErrNoRows
}

func (p produtoRepository) Create(prod models.Produto) error {
	return p.db.QueryRow("INSERT INTO produto(nome, descricao, stock, preco) VALUES($1, $2, $3, $4) RETURNING id",
		prod.Nome, prod.Descricao, prod.Stock, prod.Preco).Scan(&prod.Id)
}

func (p produtoRepository) Update(prod models.Produto) error {
	return p.db.QueryRow("UPDATE produto SET name=$1, descricao=$2, stock=$3, preco=$4, ativo=$5, data_atualizacao=$6  WHERE id=$7",
		prod.Nome, prod.Descricao, prod.Stock, prod.Preco, prod.Ativo, prod.DataAtualizacao).Scan(&prod.Id)
}

func (p produtoRepository) Delete(id int) error {
	return p.db.QueryRow("DELETE FROM produto WHERE id=$1", id).Scan(&id)
}
