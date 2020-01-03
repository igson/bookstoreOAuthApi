package token_acesso

import "github.com/igson/bookstoreOAuthApi/src/utils/erros"

import "strings"


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
	
	tokenId := strings.TrimSpace(tokenAcessoId)
	
	if len(tokenId) == 0 {
		return nil, erros.MsgBadRequestErro("Token id inválido")
	}

	tokenAcesso, erro := s.repository.BuscarPorId(tokenAcessoId)
	
	if erro != nil {
		return nil, erro
	}

	return tokenAcesso, nil

}
