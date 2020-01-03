package token_acesso

import "github.com/igson/bookstoreOAuthApi/src/utils/erros"


type Repository interface {
	BuscarPorId(tokenAcessoId string) (*AccessToken, *erros.MsgErroApi)
}

type Service interface {
	BuscarPorId(tokenAcessoId string) (*AccessToken, *erros.MsgErroApi)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) BuscarPorId(tokenAcessoId string) (*AccessToken, *erros.MsgErroApi) {
	tokenAcesso, erro := s.repository.BuscarPorId(tokenAcessoId)
	
	if erro != nil {
		return nil, erro
	}

	return tokenAcesso, nil

}
