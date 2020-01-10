package domain

import (
	"fmt"
	"strings"
	"time"
	"github.com/igson/bookstoreOAuthApi/src/utils/erros"
)

const (
	tempoExpiracao = 24
)

type TokenAcesso struct {
	AccessToken  string `json:"access_token"`
	UserId       int64  `json:"user_id"`
	ClienteId    int64  `json:"cliente_id"`
	ExpiredToken int64  `json:"expired_token"`
}

func GeraTokenDeAcesso() TokenAcesso {
	return TokenAcesso{
		ExpiredToken: time.Now().UTC().Add(tempoExpiracao * time.Hour).Unix(),
	}
}

func (token *TokenAcesso) IsTokenExpirado() bool {
	now := time.Now().UTC()
	tempoExpiracao := time.Unix(token.ExpiredToken, 0)
	fmt.Println(tempoExpiracao)
	return now.After(tempoExpiracao)
}

//Validar realiza a validação dos campos
func (token *TokenAcesso) Validar() *erros.MsgErroApi {
	
	token.AccessToken = strings.TrimSpace(token.AccessToken)

	if token.AccessToken == "" {
		return erros.MsgBadRequestErro("Token id inválido")
	}
	
	if token.UserId <= 0 {
		return erros.MsgBadRequestErro("Usuário Id inválido")
	}

	if token.ClienteId <= 0 {
		return erros.MsgBadRequestErro("Cliente Id iválido")
	}

	if token.ExpiredToken <= 0 {
		return erros.MsgBadRequestErro("Token expirado")
	}

	return nil
	
}