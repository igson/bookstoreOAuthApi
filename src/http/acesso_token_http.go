package http

import "github.com/igson/bookstoreOAuthApi/src/domain/token_acesso"

import "github.com/igson/bookstoreUserApi/utils/erros"

import "github.com/gin-gonic/gin"

import "net/http"

type AccessTokenHandler interface {
	BuscarPorId(*gin.Context)
}

type accessTokenHandler struct {
	service token_acesso.Service
}

func NewHandler(service token_acesso.Service) AccessTokenHandler {
	return &accessTokenHandler {
		service: service,
	}
}

func (h *accessTokenHandler) BuscarPorId(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, "Implemente-me")
}