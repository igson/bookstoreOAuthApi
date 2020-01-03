package token_acesso

import "github.com/igson/bookstoreUserApi/utils/erros"

type Repository interface {
	BuscarPorId(string) (*AccessToken, *erros.MsgErro)
}

type Service interface {
	BuscarPorId(string) (*AccessToken, *erros.MsgErro)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) BuscarPorId(string) (*AccessToken, *erros.MsgErro) {
	return nil, nil	
}
