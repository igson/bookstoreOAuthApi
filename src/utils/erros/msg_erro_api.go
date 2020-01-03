package erros

import (
	"github.com/igson/bookstoreUserApi/domain/usuario"
	"net/http"
	"net/url"
)

type error interface {
	Error() string
	MsgBadRequestErro(mensagem string) *MsgErro
	Validar(usuario.Usuario, url.Values) *MsgErro
}

type MsgErro struct {
	Mensagem string `json:"mensagem,omitempty"`
	Status   int    `json:"status,omitempty"`
	Erro     string `json:"erro,omitempty"`
}

type erroValidacao map[string]interface{}

type classe interface{}

type MsgErroValidacao struct {
	ValidaErro erroValidacao `json:"Erro_Validação,omitempty"`
}

func (erro *MsgErro) MsgBadRequestErro(mensagem string) *MsgErro {
	return &MsgErro{
		Mensagem: mensagem, Status: http.StatusBadRequest, Erro: "BAD_REQUEST",
	}
}

func (erro *MsgErroValidacao) Validar(objeto classe, erroValidacao MsgErroValidacao) *MsgErroValidacao {
	return &MsgErroValidacao{
		ValidaErro: map[string]interface{}{"Validação_Erro": erroValidacao},
	}
}
