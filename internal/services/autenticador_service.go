package services

import (
	"fmt"
	"mini-loja/internal/dto"
	"mini-loja/internal/dto/usuario"
	"mini-loja/internal/repositories/interfaces"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type autenticadorService struct {
	_usuarioRepository interfaces.IUsuarioRepository
}

func NewAutenticadorService(usuariorepository interfaces.IUsuarioRepository) autenticadorService {
	return autenticadorService{
		_usuarioRepository: usuariorepository,
	}
}

func (a autenticadorService) AutenticarUsuario(dados usuario.UsuarioAutenticarDto) dto.ResponseApiDto {

	var salvaSenha = dados.Senha
	_senha, err := HashPassword(dados.Senha)
	if err != nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    err.Error(),
		}
	}
	dados.Senha = _senha

	retortnoUsuario, err := a._usuarioRepository.BuscarUsuarioPorEmail_Senha(dados)
	if err != nil {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Erro de sistema",
		}
	} else if retortnoUsuario.Id == 0 {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Usuário ou senha invalidos",
		}
	}

	if !validaPasswordHash(salvaSenha, retortnoUsuario.Senha) {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    "Senha inválida",
		}
	}

	//Gerar token de usuário
	token, err := GerarToken(retortnoUsuario.Id, retortnoUsuario.Email)
	if err != nil {
		panic(err)
	}
	//fmt.Println("Token JWT:", token)

	// Validando...
	validado, err := ValidarToken(token)
	if err != nil || !validado.Valid {
		return dto.ResponseApiDto{
			Status: false,
			Msg:    fmt.Sprintf("Erro de validação do Token: %v", err),
		}
	}

	//claims := validado.Claims.(jwt.MapClaims)
	//fmt.Println("User ID:", claims["user_id"])
	//fmt.Println("Email:", claims["email"])

	usuarioAutenticado := usuario.UsuarioAutenticadoDto{
		Token: token,
	}

	return dto.ResponseApiDto{
		Status: true,
		Msg:    "Dados do usuário autenticados com sucesso",
		Data:   usuarioAutenticado,
	}
}

var jwtSecret = []byte("sua-chave-secreta-super-segura")

// Gera um token JWT com informações do usuário
func GerarToken(userID int, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // expira em 24h
		"iat":     time.Now().Unix(),                     // emitido em
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Valida e extrai dados do token
func ValidarToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Garante que está usando o algoritmo certo
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inválido")
		}
		return jwtSecret, nil
	})
}

// Uteis

func validaPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
