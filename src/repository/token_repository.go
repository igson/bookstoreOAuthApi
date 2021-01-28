package repository

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/igson/bookstoreOAuthApi/src/databases/cassandra"
	"github.com/igson/bookstoreOAuthApi/src/domain"
	"github.com/igson/bookstoreOAuthApi/src/utils/errors"
)

const (
	queryGetAccessToken         = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token = ?"
	queryCreateAccessToken      = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires          = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?"
)

type TokenRepository interface {
	GetByID(string) (*domain.Token, *errors.RestErroAPI)
	Create(domain.Token) *errors.RestErroAPI
	UpdateExpirationTime(domain.Token) *errors.RestErroAPI
}

type tokenRepository struct {
}

func NewTokenRepository() TokenRepository {
	return &tokenRepository{}
}

func (r *tokenRepository) GetByID(tokenID string) (*domain.Token, *errors.RestErroAPI) {

	var token domain.Token

	if erro := cassandra.GetSession().Query(queryGetAccessToken, tokenID).Scan(&token.AccessToken, &token.AccessToken, &token.UserID, &token.ClientID, &token.Expires); erro != nil {
		fmt.Println(erro)
		if erro == gocql.ErrNotFound {
			return nil, errors.NewNotFoundErro("Token não encontrado")
		}
		return nil, errors.NewInternalServerError(erro.Error())
	}

	return &token, nil

}

func (r *tokenRepository) Create(token domain.Token) *errors.RestErroAPI {

	if erro := cassandra.GetSession().Query(queryCreateAccessToken,
		token.AccessToken,
		token.UserID,
		token.ClientID,
		token.Expires,
	).Exec(); erro != nil {
		return errors.NewInternalServerError("Erro ao tentar cadastrar o token")
	}

	return nil
}

func (r *tokenRepository) UpdateExpirationTime(token domain.Token) *errors.RestErroAPI {
	if err := cassandra.GetSession().Query(queryUpdateExpires,
		token.Expires,
		token.AccessToken,
	).Exec(); err != nil {
		return errors.NewInternalServerError("Erro ao tentar atualizar o token")
	}
	return nil
}
