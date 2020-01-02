package token_acesso

import "time"

const (
	tempoExpiracao = 24
)

type AccessToken struct {
	AccessToken	string `json:"access_token"`
	UserId int64 `json:"user_id"`
	ClienteId int64 `json:"cliente_id"`
	ExpiredToken  int64 `json:"expired_token"`
}


func GeraTokenDeAcesso() AccessToken {
	return AccessToken{
		ExpiredToken: time.Now().UTC().Add(tempoExpiracao * time.Hour).Unix(),
	}
}

func (token AccessToken) IsTokenExpirado() bool {
	return false
} 