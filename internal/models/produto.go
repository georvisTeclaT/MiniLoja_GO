package models

import "time"

type Produto struct {
	Id              int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Nome            string    `json:"nome"`
	Descricao       string    `json:"descricao"`
	Stock           int       `json:"stock"`
	Preco           float64   `json:"preco"`
	Ativo           bool      `json:"ativo"`
	DataCriacao     time.Time `json:"data_criacao"`
	DataAtualizacao time.Time `json:"data_atualizacao"`
}
