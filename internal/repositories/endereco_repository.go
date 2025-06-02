package repositories

import (
	"database/sql"
	enderecousuario "mini-loja/internal/dto/endereco_usuario"
	"mini-loja/internal/models"
	"mini-loja/internal/repositories/interfaces"
	"time"
)

type enderecoRepository struct {
	db *sql.DB
}

func NewEnderecoRepository(db *sql.DB) interfaces.IEnderecoRepository {
	return enderecoRepository{db: db}
}

func (e enderecoRepository) GetAll() ([]enderecousuario.EnderecoUsuarioDto, error) {

	rows, err := e.db.Query("SELECT id, usuario_id, nome_rua, numero, complemento, bairro, cidade, estado, cep, data_criacao  FROM endereco_usuario")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enderecosList []enderecousuario.EnderecoUsuarioDto
	for rows.Next() {
		var endereco enderecousuario.EnderecoUsuarioDto
		var dataCriacao time.Time
		if err := rows.Scan(
			&endereco.Id,
			&endereco.Usuario,
			&endereco.NomeRua,
			&endereco.Numero,
			&endereco.Complemento,
			&endereco.Bairro,
			&endereco.Cidade,
			&endereco.Estado,
			&endereco.Cep,
			&dataCriacao); err != nil {
			return nil, err
		}

		// Formata a data e coloca no campo string
		endereco.DataCriacao = dataCriacao.Format("02/01/2006")

		enderecosList = append(enderecosList, endereco)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return enderecosList, nil
}

func (e enderecoRepository) GetById(id int) (enderecousuario.EnderecoUsuarioDto, error) {

	rows, err := e.db.Query("SELECT id, usuario_id, nome_rua, numero, complemento, bairro, cidade, estado, cep, data_criacao  FROM endereco_usuario WHERE id = $1", id)
	if err != nil {
		return enderecousuario.EnderecoUsuarioDto{}, err
	}
	defer rows.Close()

	var endereco enderecousuario.EnderecoUsuarioDto
	if rows.Next() {
		var dataCriacao time.Time
		if err := rows.Scan(
			&endereco.Id,
			&endereco.Usuario,
			&endereco.NomeRua,
			&endereco.Numero,
			&endereco.Complemento,
			&endereco.Bairro,
			&endereco.Cidade,
			&endereco.Estado,
			&endereco.Cep,
			&dataCriacao); err != nil {
			return endereco, err
		}

		// Formata a data e coloca no campo string
		endereco.DataCriacao = dataCriacao.Format("02/01/2006")
	}

	return endereco, sql.ErrNoRows
}

func (e enderecoRepository) GetEnderecoById(idEndereco int, idUsuario int) (models.EnderecoUsuario, error) {

	rows, err := e.db.Query("SELECT id, usuario_id, nome_rua, numero, complemento, bairro, cidade, estado, cep, data_criacao, data_atualizacao FROM endereco_usuario WHERE id = $1 and usuario_id = $2", idEndereco, idUsuario)
	if err != nil {
		return models.EnderecoUsuario{}, err
	}
	defer rows.Close()

	var endereco models.EnderecoUsuario
	if rows.Next() {
		if err := rows.Scan(
			&endereco.Id,
			&endereco.UsuarioId,
			&endereco.NomeRua,
			&endereco.Numero,
			&endereco.Complemento,
			&endereco.Bairro,
			&endereco.Cidade,
			&endereco.Estado,
			&endereco.Cep,
			&endereco.DataCriacao,
			&endereco.DataAtualizacao); err != nil {
			return endereco, err
		}
	}

	return endereco, sql.ErrNoRows
}

func (e enderecoRepository) Create(endereco models.EnderecoUsuario) error {
	return e.db.QueryRow("INSERT INTO endereco_usuario (usuario_id, nome_rua, numero, complemento, bairro, cidade, estado, cep) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		endereco.UsuarioId, endereco.NomeRua, endereco.Numero, endereco.Complemento, endereco.Bairro, endereco.Cidade, endereco.Estado, endereco.Cep).Scan(&endereco.Id)
}

func (e enderecoRepository) Update(endereco models.EnderecoUsuario) error {
	return e.db.QueryRow("UPDATE endereco_usuario SET nome_rua=$1, numero=$2, complemento=$3, bairro=$4, cidade=$5, estado=$6, cep=$7, data_atualizacao=$8 WHERE id=$9",
		endereco.NomeRua, endereco.Numero, endereco.Complemento, endereco.Bairro, endereco.Cidade, endereco.Estado, endereco.Cep, endereco.DataAtualizacao).Scan(&endereco.Id)
}

func (e enderecoRepository) Delete(id int) error {
	return e.db.QueryRow("DELETE FROM endereco_usuario WHERE id=$1", id).Scan(&id)
}
