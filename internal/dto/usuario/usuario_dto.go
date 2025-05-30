package usuario

type UsuarioDto struct {
	Id          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Nome        string `json:"nome"`
	Sobrenome   string `json:"sobrenome"`
	Email       string `json:"email"`
	Telefone    string `json:"telefone"`
	Ativo       bool   `json:"ativo"`
	DataCriacao string `json:"data_criacao"`
}
