package client

import (
	"encoding/json"
	"time"

	"github.com/igson/bookstoreOAuthApi/src/rest/dto"
	"github.com/igson/bookstoreOAuthApi/src/utils/erros"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usuarioClient = rest.RequestBuilder{
		BaseURL: "localhost:8080/api/usuarios",
		Timeout: 100 * time.Millisecond,
	}
)

type UsuarioRestClient interface {
	LoginUsuario(string, string) (*dto.UsuarioDTO, *erros.MsgErroApi)
}

type usuarioRestClient struct{}

func NewUsuarioRestClient() UsuarioRestClient {
	return &usuarioRestClient{}
}

func (u *usuarioRestClient) LoginUsuario(email string, senha string) (*dto.UsuarioDTO, *erros.MsgErroApi) {
	usuarioDTO := dto.UsuarioDTO{
		Email: email,
		Senha: senha,
	}

	resposta := usuarioClient.Post("/login", usuarioDTO)

	if resposta == nil || resposta.Response == nil {
		return nil, erros.MsgInternalServerError("Resposta inválida ao tentar efetuar login de usuário")
	}

	if resposta.StatusCode > 299 {
		var restErro erros.MsgErroApi
		erro := json.Unmarshal(resposta.Bytes(), &restErro)
		if erro != nil {
			return nil, erros.MsgInternalServerError("Erro na interface ao tentar efetuar login de usuário")
		}
		return nil, &restErro
	}

	var usuario dto.UsuarioDTO

	if erro := json.Unmarshal(resposta.Bytes(), &usuario); erro != nil {
		return nil, erros.MsgInternalServerError("erro ao tentar desserializar resposta")
	}

	return &usuario, nil
}
