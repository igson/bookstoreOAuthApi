package service

import (
	"strings"

	"github.com/igson/bookstoreOAuthApi/src/domain"
	"github.com/igson/bookstoreOAuthApi/src/repository"
	"github.com/igson/bookstoreOAuthApi/src/repository/rest"
	"github.com/igson/bookstoreOAuthApi/src/utils/errors"
)

type TokenService interface {
	GetByID(string) (*domain.Token, *errors.RestErroAPI)
	Create(requestToken domain.AccessTokenRequest) (*domain.Token, *errors.RestErroAPI)
	UpdateExpirationTime(domain.Token) *errors.RestErroAPI
}

type tokenService struct {
	restUsersRepostory rest.RestUserRepository
	toTokenRepository  repository.TokenRepository
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

func (s *tokenService) Create(requestToken domain.AccessTokenRequest) (*domain.Token, *errors.RestErroAPI) {

	if err := requestToken.Validate(); err != nil {
		return nil, err
	}

	//TODO: Support both grant types: client_credentials and password

	// Authenticate the user against the Users API:
	user, err := s.restUsersRepostory.LoginUser(requestToken.Username, requestToken.Password)

	if err != nil {
		return nil, err
	}

	at := domain.GetNewAccessToken(user.ID)

	at.Generate()

	if err := s.toTokenRepository.Create(at); err != nil {
		return nil, err
	}

	return &at, nil

}

func (s *tokenService) UpdateExpirationTime(token domain.Token) *errors.RestErroAPI {

	if erro := token.Validate(); erro != nil {
		return erro
	}

	return s.toTokenRepository.UpdateExpirationTime(token)

}
