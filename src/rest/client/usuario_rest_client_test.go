package client

import (
	"testing"
	"os"
	"github.com/mercadolibre/golang-restclient/rest"
	"net/http"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}


func TestLoginUsuarioTimeout(t *testing.T) {

	rest.FlushMockups()
	
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost:8080/api/usuarios/login",
		ReqBody:      `{"email":"igson@gmail.com","password":"tce123"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})

	usuarioClient := usuarioRestClient{}

	usuario, erro := usuarioClient.LoginUsuario("igson@gmail.com", "tce123")

	assert.Nil(t, usuario)
	assert.NotNil(t, erro)
	assert.EqualValues(t, http.StatusInternalServerError, erro.Status)
	assert.EqualValues(t, "invalid restclient response when trying to login user", erro.Mensagem)
}

func TestLoginUsuarioInterfaceInvalida(t *testing.T) {

}

func TestLoginUsuarioCredencialInvalida(t *testing.T) {

}

func TestLoginUsuarioRespostaInvalida(t *testing.T) {

}

func TestLoginComSucesso(t *testing.T) {

}
