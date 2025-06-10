package usuario

type UsuarioAutenticarDto struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}
