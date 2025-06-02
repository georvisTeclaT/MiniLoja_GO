package enderecousuario

type EnderecoUsuarioAddUpdateDto struct {
	NomeRua     string `json:"nome_rua"`
	Numero      string `json:"numero"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Cidade      string `json:"cidade"`
	Estado      string `json:"estado"`
	Cep         string `json:"cep"`
}
