package models

import "time"

type Usuario struct {
	Id              int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Nome            string    `json:"nome"`
	Sobrenome       string    `json:"sobrenome"`
	DataCadastro    time.Time `json:"data_cadastro"`
	DataAtualizacao time.Time `json:"data_atualizacao"`
}
