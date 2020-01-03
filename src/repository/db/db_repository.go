package db

import (
	"github.com/igson/bookstoreOAuthApi/src/utils/erros"
	"github.com/igson/bookstoreOAuthApi/src/domain/token_acesso"
	
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type dbRepository struct {}

type DbRepository interface {
	BuscarPorId(tokenAcessoId string) (*token_acesso.AccessToken, *erros.MsgErroApi)
}

func (r *dbRepository) BuscarPorId(tokenAcessoId string) (*token_acesso.AccessToken, *erros.MsgErroApi) {
	return nil, erros.MsgInternalServerError("Erro de conexão com o banco de dados")
}