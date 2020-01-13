package app

import (
	"github.com/gin-gonic/gin"
	"github.com/igson/bookstoreOAuthApi/src/controller"
	"github.com/igson/bookstoreOAuthApi/src/repository"
	"github.com/igson/bookstoreOAuthApi/src/service"
)

var (
	rota = gin.Default()
)

//StartApplication - dá início a aplicação
func StartApplication() {

	atHandler := controller.NewHandler(service.NewTokenService(repository.NewTokenRepository()))

	token := rota.Group("/api")
	{
		token.GET("/oauth/tokens/:acessoTokenId", atHandler.BuscarPorId)
		token.POST("/oauth/tokens", atHandler.CriarTokenAcesso)
	}

	rota.Run(":8080")

}
