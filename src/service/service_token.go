package service

import (
	"github.com/igson/bookstoreOAuthApi/src/domain"
	"github.com/igson/bookstoreOAuthApi/src/repository"
	"strings"
	"github.com/igson/bookstoreOAuthApi/src/utils/erros"
)

type TokenService interface {
	CriarTokenAcesso(domain.TokenAcesso) *erros.MsgErroApi
	AtualizarTokenExpirado(domain.TokenAcesso) *erros.MsgErroApi
	BuscarPorId(string) (*domain.TokenAcesso, *erros.MsgErroApi)
}

type tokenService struct {
	tokenRepository repository.TokenRepository
}

//NewTokenService injeta a dependência do cassandra service
func NewTokenService(repo repository.TokenRepository) TokenService {
	return &tokenService{
		tokenRepository: repo,
	}
}

func (service *tokenService) CriarTokenAcesso(token domain.TokenAcesso) *erros.MsgErroApi {

	if erro := token.Validar(); erro != nil {
		return erro
	}

	return service.tokenRepository.CriarTokenAcesso(token)
}

func (service *tokenService) AtualizarTokenExpirado(token domain.TokenAcesso) *erros.MsgErroApi {

	if erro := token.Validar(); erro != nil {
		return erro
	} 

	return service.tokenRepository.AtualizarTokenExpirado(token)
}

func (service *tokenService) BuscarPorId(tokenAcessoId string) (*domain.TokenAcesso, *erros.MsgErroApi) {

	tokenId := strings.TrimSpace(tokenAcessoId)

	if len(tokenId) == 0 {
		return nil, erros.MsgBadRequestErro("Token id inválido")
	}

	token, erro := service.tokenRepository.BuscarPorId(tokenId)

	if erro != nil {
		return nil, erro
	}

	return token, nil

}
