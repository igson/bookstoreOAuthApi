package app

import (
	"github.com/gin-gonic/gin"
	"github.com/igson/bookstoreOAuthApi/src/controllers"
	"github.com/igson/bookstoreOAuthApi/src/repository"
	"github.com/igson/bookstoreOAuthApi/src/service"
)

var (
	rota = gin.Default()
)

//StartApplication - dá início a aplicação
func StartApplication() {

	tokenController := controllers.NewTokenController(service.NewTokenService(repository.NewTokenRepository()))

	token := rota.Group("")
	{
		token.GET("/oauth/tokens/:acessoTokenId", tokenController.GetByID)
		token.POST("/oauth/tokens", tokenController.Create)
	}

	rota.Run(":8080")

}
