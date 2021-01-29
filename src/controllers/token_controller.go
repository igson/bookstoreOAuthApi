package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igson/bookstoreOAuthApi/src/domain"
	"github.com/igson/bookstoreOAuthApi/src/service"
	"github.com/igson/bookstoreOAuthApi/src/utils/errors"
)

type TokenController interface {
	GetByID(*gin.Context)
	Create(c *gin.Context)
}

type tokenController struct {
	toTokenService service.TokenService
}

func NewTokenController(tokenService service.TokenService) TokenController {
	return &tokenController{
		toTokenService: tokenService,
	}
}

func (tc *tokenController) GetByID(c *gin.Context) {

	token, erro := tc.toTokenService.GetByID(c.Param("acessoTokenId"))

	if erro != nil {
		c.JSON(erro.StatusCode, erro)
		return
	}

	c.JSON(http.StatusOK, token)
}

func (tc *tokenController) Create(c *gin.Context) {

	var at domain.AccessTokenRequest

	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("invalid json error body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	accessToken, err := tc.toTokenService.Create(at)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusCreated, accessToken)

}
