package service

import (
	"strings"

	"github.com/igson/bookstoreOAuthApi/src/domain"
	"github.com/igson/bookstoreOAuthApi/src/repository"
	"github.com/igson/bookstoreOAuthApi/src/utils/errors"
)

type TokenService interface {
	GetByID(string) (*domain.Token, *errors.RestErroAPI)
	Create(domain.Token) *errors.RestErroAPI
	UpdateExpirationTime(domain.Token) *errors.RestErroAPI
}

type tokenService struct {
	toTokenRepository repository.TokenRepository
}

//NewTokenService injeção
func NewTokenService(tokenRepository repository.TokenRepository) TokenService {
	return &tokenService{
		toTokenRepository: tokenRepository,
	}
}

func (s *tokenService) GetByID(tokenID string) (*domain.Token, *errors.RestErroAPI) {

	tokenID = strings.TrimSpace(tokenID)

	if len(tokenID) == 0 {
		return nil, errors.NewBadRequestError("Token invalido")
	}

	token, erro := s.toTokenRepository.GetByID(tokenID)

	if erro != nil {
		return nil, erro
	}

	return token, nil

}

func (s *tokenService) Create(token domain.Token) *errors.RestErroAPI {

	if erro := token.Validate(); erro != nil {
		return erro
	}

	return s.toTokenRepository.Create(token)

}

func (s *tokenService) UpdateExpirationTime(token domain.Token) *errors.RestErroAPI {

	if erro := token.Validate(); erro != nil {
		return erro
	}

	return s.toTokenRepository.UpdateExpirationTime(token)

}
