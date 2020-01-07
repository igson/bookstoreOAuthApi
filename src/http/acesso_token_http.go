package http

import(
	"github.com/igson/bookstoreOAuthApi/src/domain/token_acesso"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//FIXME melhorar o código
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
	
	tokenAcesso, erro := h.service.BuscarPorId(strings.TrimSpace(ctx.Param("acessoTokenId")))

	if erro != nil {
		ctx.JSON(erro.Status, erro)
		return 
	}

	ctx.JSON(http.StatusOK, tokenAcesso)
}