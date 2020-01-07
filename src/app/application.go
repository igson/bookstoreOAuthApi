package app

import (
	"github.com/igson/bookstoreOAuthApi/src/repository"
	"github.com/gin-gonic/gin"
	"github.com/igson/bookstoreOAuthApi/src/clients/cassandra"
	"github.com/igson/bookstoreOAuthApi/src/http"
	"github.com/igson/bookstoreOAuthApi/src/service"
)

var (
	rota = gin.Default()
)

func StartApplication() {

	sessao, dbErro := cassandra.GetSession()

	if dbErro != nil {
		panic(dbErro)
	}

	sessao.Close()

	atHandler := http.NewHandler(service.NewTokenService(repository.NewTokenRepository()))

	token := rota.Group("/api")
	{
		token.GET("/oauth/tokens/:acessoTokenId", atHandler.BuscarPorId)
		token.POST("/oauth/tokens", atHandler.CriarTokenAcesso)
	}

	rota.Run(":8080")

}
