package domain

import (
	"fmt"
	"strings"
	"time"

	"github.com/igson/bookstoreOAuthApi/src/utils/crypto_utils"
	"github.com/igson/bookstoreOAuthApi/src/utils/erros"
)

const (
	tempoExpiracao             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type TokenAcesso struct {
	AccessToken  string `json:"access_token"`
	UserId       int64  `json:"user_id"`
	ClienteId    int64  `json:"cliente_id"`
	ExpiredToken int64  `json:"expired_token"`
}

type RequestTokenAcesso struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	Nome         string `json:"nome"`
	Senha        string `json:"senha"`
	ClientId     int64  `json:"client_id"`
	ClientSecret int64  `json:"client_secret"`
}

func (token *RequestTokenAcesso) Validar() *erros.MsgErroApi {
	switch token.GrantType {
	case grantTypePassword:
		break

	case grantTypeClientCredentials:
		break

	default:
		return erros.MsgBadRequestErro("invalid grant_type parameter")
	}
	return nil

}

func GeraTokenDeAcesso() TokenAcesso {
	return TokenAcesso{
		ExpiredToken: time.Now().UTC().Add(tempoExpiracao * time.Hour).Unix(),
	}
}

func GetNewAccessToken(userId int64) TokenAcesso {
	return TokenAcesso{
		UserId:  userId,
		ExpiredToken: time.Now().UTC().Add(tempoExpiracao * time.Hour).Unix(),
	}
}

func (at *TokenAcesso) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.ExpiredToken))
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
