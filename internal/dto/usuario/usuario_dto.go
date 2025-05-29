package usuario

type UsuarioDto struct {
	Id           int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Nome         string `json:"nome"`
	Sobrenome    string `json:"sobrenome"`
	DataCadastro string `json:"data_cadastro"`
}
