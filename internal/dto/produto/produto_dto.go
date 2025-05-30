package produto

type ProdutoDto struct {
	Id          int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Nome        string  `json:"nome"`
	Descricao   string  `json:"descricao"`
	Stock       int     `json:"stock"`
	Preco       float64 `json:"preco"`
	Ativo       bool    `json:"ativo"`
	DataCriacao string  `json:"data_criacao"`
}
