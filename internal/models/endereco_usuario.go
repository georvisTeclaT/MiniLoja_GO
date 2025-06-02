package models

import "time"

type EnderecoUsuario struct {
	Id              int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UsuarioId       int       `json:"usuarioId"`
	NomeRua         string    `json:"nome_rua"`
	Numero          string    `json:"numero"`
	Complemento     string    `json:"complemento"`
	Bairro          string    `json:"bairro"`
	Cidade          string    `json:"cidade"`
	Estado          string    `json:"estado"`
	Cep             string    `json:"cep"`
	DataCriacao     time.Time `json:"data_criacao"`
	DataAtualizacao time.Time `json:"data_atualizacao"`
}
