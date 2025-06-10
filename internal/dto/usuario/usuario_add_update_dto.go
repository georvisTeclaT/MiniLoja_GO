package usuario

type UsuarioAddUpdateDto struct {
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Email     string `json:"email"`
	Telefone  string `json:"telefone"`
	Ativo     bool   `json:"ativo"`
	Senha     string `json:"senha"`
}
