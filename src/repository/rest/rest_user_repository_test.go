package rest

import (
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http:localhost:8082/users/login",
		HTTPMethod:   http.MethodPost,
		ReqHeaders:   map[string][]string{},
		ReqBody:      `{"email":"igson@gmail.com","password":"tce1234"}`,
		RespHTTPCode: -1,
		RespHeaders:  map[string][]string{},
		RespBody:     `{}`,
	})

	repository := userRepository{}

	user, err := repository.LoginUser("igson@gmail.com", "tce123")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid restclient response when trying to login user", err.Message)
}
