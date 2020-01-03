package app

import (
	"github.com/igson/bookstoreOAuthApi/src/domain/token_acesso"
	"github.com/gin-gonic/gin"
	"github.com/igson/bookstoreOAuthApi/src/repository/db"
	"github.com/igson/bookstoreOAuthApi/src/http"
)

var (
	rotas = gin.Default()
)

func StartApplication() {
	dbRepository := db.NewRepository()
	tokenService := token_acesso.NewService(dbRepository)
	atHandler := http.NewHandler(tokenService)	

	rotas.GET("/oauth/acesso_token/:acessoTokenId", atHandler.BuscarPorId)

	rotas.Run(":8080")
}