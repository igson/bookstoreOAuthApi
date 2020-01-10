package http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/igson/bookstoreOAuthApi/src/domain"
	"github.com/igson/bookstoreOAuthApi/src/service"
	"github.com/igson/bookstoreOAuthApi/src/utils/erros"
)

//AccessTokenHandler interface de acesso ao controller
type AccessTokenHandler interface {
	BuscarPorId(*gin.Context)
	CriarTokenAcesso(ctx *gin.Context)
}

type accessTokenHandler struct {
	tokenService service.TokenService
}

//NewHandler realiza o IOC de acesso ao controller
func NewHandler(service service.TokenService) AccessTokenHandler {
	return &accessTokenHandler{
		tokenService: service,
	}
}

//BuscarPorId realiza a busca por ID do usuário
func (h *accessTokenHandler) BuscarPorId(ctx *gin.Context) {

	fmt.Println("ID do Token: ", strings.TrimSpace(ctx.Param("acessoTokenId")))

	tokenAcesso, erro := h.tokenService.BuscarPorId(strings.TrimSpace(ctx.Param("acessoTokenId")))

	if erro != nil {
		ctx.JSON(erro.Status, erro)
		return
	}

	ctx.JSON(http.StatusOK, tokenAcesso)
}

//CriarTokenAcesso reponsável por criar o token de acesso
func (h *accessTokenHandler) CriarTokenAcesso(ctx *gin.Context) {

	var token domain.TokenAcesso

	if jsonErroBind := ctx.ShouldBindJSON(&token); jsonErroBind != nil {
		msgErro := erros.MsgBadRequestErro("Formato de campos inválido.")
		ctx.JSON(msgErro.Status, msgErro)
		return
	}

	if msgErro := h.tokenService.CriarTokenAcesso(token); msgErro != nil {
		ctx.JSON(msgErro.Status, msgErro)
		return
	}

	ctx.JSON(http.StatusCreated, token )

}
