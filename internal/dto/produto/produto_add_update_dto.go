package produto

type ProdutoAddUpdateDto struct {
	Nome      string  `json:"nome"`
	Descricao string  `json:"descricao"`
	Stock     int     `json:"stock"`
	Preco     float64 `json:"preco"`
}
