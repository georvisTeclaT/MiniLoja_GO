package repositories

import (
	"database/sql"
	"mini-loja/internal/dto/usuario"
	"mini-loja/internal/models"
	"mini-loja/internal/repositories/interfaces"
	"time"
)

type usuarioRepository struct {
	db *sql.DB
}

func NewUsuarioRepository(db *sql.DB) interfaces.IUsuarioRepository {
	return usuarioRepository{db: db}
}

func (u usuarioRepository) GetAll() ([]usuario.UsuarioDto, error) {

	rows, err := u.db.Query("SELECT id, nome, sobrenome, data_cadastro FROM usuario")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuariosList []usuario.UsuarioDto
	for rows.Next() {
		var usuario usuario.UsuarioDto
		var dataCadastro time.Time
		if err := rows.Scan(&usuario.Id, &usuario.Nome, &usuario.Sobrenome, &dataCadastro); err != nil {
			return nil, err
		}

		// Formata a data e coloca no campo string
		usuario.DataCadastro = dataCadastro.Format("02/01/2006")

		usuariosList = append(usuariosList, usuario)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return usuariosList, nil
}

func (u usuarioRepository) GetByID(id int) (usuario.UsuarioDto, error) {
	rows, err := u.db.Query("SELECT id, nome, sobrenome, data_cadastro FROM usuario WHERE id = $1", id)
	if err != nil {
		return usuario.UsuarioDto{}, err
	}
	defer rows.Close()

	var usuario usuario.UsuarioDto
	if rows.Next() {
		var dataCadastro time.Time
		if err := rows.Scan(&usuario.Id, &usuario.Nome, &usuario.Sobrenome, &dataCadastro); err != nil {
			return usuario, err
		}

		// Formata a data e coloca no campo string
		usuario.DataCadastro = dataCadastro.Format("02/01/2006")
	}

	return usuario, sql.ErrNoRows
}

func (u usuarioRepository) GetUsuarioByID(id int) (models.Usuario, error) {
	rows, err := u.db.Query("SELECT id, nome, sobrenome, data_cadastro, data_atualizacao FROM usuario WHERE id = $1", id)
	if err != nil {
		return models.Usuario{}, err
	}
	defer rows.Close()

	var usuario models.Usuario
	if rows.Next() {
		if err := rows.Scan(&usuario.Id, &usuario.Nome, &usuario.Sobrenome, &usuario.DataCadastro, usuario.DataAtualizacao); err != nil {
			return usuario, err
		}
		return usuario, nil
	}

	return usuario, sql.ErrNoRows
}

func (u usuarioRepository) Create(usuario models.Usuario) error {
	return u.db.QueryRow("INSERT INTO usuario(nome, sobrenome) VALUES($1, $2) RETURNING id", usuario.Nome, usuario.Sobrenome).Scan(&usuario.Id)
}

func (u usuarioRepository) Update(usuario models.Usuario) error {
	return u.db.QueryRow("UPDATE usuario SET nome=$1, sobrenome=$2, data_atualizacao=$3 WHERE id=$4", usuario.Nome, usuario.Sobrenome, usuario.DataAtualizacao, usuario.Id).Scan(&usuario.Id)
}

func (u usuarioRepository) Delete(id int) error {
	return u.db.QueryRow("DELETE FROM usuario WHERE id = $1", id).Scan(&id)
}
