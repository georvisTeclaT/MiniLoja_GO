package enderecousuario

type EnderecoUsuarioDto struct {
	Id          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Usuario     string `json:"usuario"`
	NomeRua     string `json:"nome_rua"`
	Numero      string `json:"numero"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Cidade      string `json:"cidade"`
	Estado      string `json:"estado"`
	Cep         string `json:"cep"`
	DataCriacao string `json:"data_criacao"`
}
