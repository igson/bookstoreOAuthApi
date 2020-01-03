package db

import (
	"github.com/igson/bookstoreOAuthApi/src/domain/token_acesso"
	"github.com/igson/bookstoreUserApi/utils/erros"
)


func NewRepository() DbRepository {
	return &dbRepository{}
}

type dbRepository struct {}

type DbRepository interface {
	BuscarPorId(string) (*token_acesso.AccessToken, *erros.MsgErro)
}

func (r *dbRepository) BuscarPorId(id string) (*token_acesso.AccessToken, *erros.MsgErro) {
	return nil, nil
}