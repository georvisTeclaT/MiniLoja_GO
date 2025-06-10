package models

import "time"

type Usuario struct {
	Id              int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Nome            string    `json:"nome"`
	Sobrenome       string    `json:"sobrenome"`
	Email           string    `json:"email"`
	Telefone        string    `json:"telefone"`
	Ativo           bool      `json:"ativo"`
	DataCriacao     time.Time `json:"data_criacao"`
	DataAtualizacao time.Time `json:"data_atualizacao"`
	Senha           string    `json:"senha"`
}
