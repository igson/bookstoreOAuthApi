package app

import (
	"github.com/igson/bookstoreOAuthApi/src/domain/token_acesso"
	"github.com/igson/bookstoreOAuthApi/src/clients/cassandra"
	"github.com/gin-gonic/gin"
	"github.com/igson/bookstoreOAuthApi/src/repository/db"
	"github.com/igson/bookstoreOAuthApi/src/http"
)

var (
	rotas = gin.Default()
)

func StartApplication() {

	sessao, dbErro := cassandra.GetSession()

	if dbErro != nil {
		panic(dbErro)
	}

	sessao.Close()

	atHandler := http.NewHandler(token_acesso.NewService(db.NewRepository()))	
	rotas.GET("/oauth/acesso_token/:acessoTokenId", atHandler.BuscarPorId)
	
	rotas.Run(":8080")
}