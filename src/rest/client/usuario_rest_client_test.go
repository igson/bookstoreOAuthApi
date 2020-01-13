package client

import (
	"fmt"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUsuarioTimeout(t *testing.T) {
	rest.FlushMockups()

	usuarioClient := usuarioRestClient{}
	rest, erro := usuarioClient.LoginUsuario("igson@gmail.com", "tce123")
	fmt.Println(rest)
	fmt.Println(erro)
}

func TestLoginUsuarioInterfaceInvalida(t *testing.T) {

}

func TestLoginUsuarioCredencialInvalida(t *testing.T) {

}

func TestLoginUsuarioRespostaInvalida(t *testing.T) {

}

func TestLoginComSucesso(t *testing.T) {

}
