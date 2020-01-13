package repository

import (
	"github.com/gocql/gocql"
	"github.com/igson/bookstoreOAuthApi/src/clients/cassandra"
	"github.com/igson/bookstoreOAuthApi/src/domain"
	"github.com/igson/bookstoreOAuthApi/src/utils/erros"
)

const (
	queryGetTokenAcesso         = "SELECT access_token, client_id, user_id, expires FROM access_tokens WHERE access_token=?;"
	queryCriarTokenAcesso       = "INSERT INTO access_tokens(access_token, client_id, user_id, expires) VALUES (?,?,?,?);"
	queryAtualizarTokenExpirado = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

func NewTokenRepository() TokenRepository {
	return &tokenRepository{}
}

type tokenRepository struct{}

//TokenRepository interface do repositorio
type TokenRepository interface {
	CriarTokenAcesso(domain.TokenAcesso) *erros.MsgErroApi
	AtualizarTokenExpirado(domain.TokenAcesso) *erros.MsgErroApi
	BuscarPorId(tokenAcessoID string) (*domain.TokenAcesso, *erros.MsgErroApi)
}

func (repo *tokenRepository) CriarTokenAcesso(token domain.TokenAcesso) *erros.MsgErroApi {

	if erro := cassandra.GetSession().Query(queryCriarTokenAcesso,
		token.AccessToken,
		token.UserId,
		token.ClienteId,
		token.ExpiredToken,
	).Exec(); erro != nil {
		return erros.MsgInternalServerError(erro.Error())
	}

	return nil

}

func (repo *tokenRepository) AtualizarTokenExpirado(token domain.TokenAcesso) *erros.MsgErroApi {

	if erro := cassandra.GetSession().Query(queryAtualizarTokenExpirado,
		token.ExpiredToken,
		token.AccessToken,
	).Exec(); erro != nil {
		return erros.MsgInternalServerError(erro.Error())
	}

	return nil

}

//BuscarPorId metodo que realiza s budsca do token pelo ID
func (repo *tokenRepository) BuscarPorId(tokenAcessoID string) (*domain.TokenAcesso, *erros.MsgErroApi) {

	var tokenAcesso domain.TokenAcesso

	if erro := cassandra.GetSession().Query(queryGetTokenAcesso, tokenAcessoID).Scan(
		&tokenAcesso.AccessToken,   //param 1
		&tokenAcesso.ClienteId,     //param 2
		&tokenAcesso.UserId,        //param 3
		&tokenAcesso.ExpiredToken); //param 4
	erro != nil {
		if erro == gocql.ErrNotFound {
			return nil, erros.MsgNotFoundErro("Nenhum token de acesso encontrado com o id fornecido")
		}
		return nil, erros.MsgInternalServerError(erro.Error())
	}

	return &tokenAcesso, nil
}
