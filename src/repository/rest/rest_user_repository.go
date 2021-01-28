package rest

import (
	"encoding/json"
	"time"

	"github.com/igson/bookstoreOAuthApi/src/utils/errors"
	"github.com/igson/bookstoreUsersApi/domain/users"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8082",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUserRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErroAPI)
}

type userRepository struct{}

func NewRestUserRepository() RestUserRepository {
	return &userRepository{}
}

func (r *userRepository) LoginUser(email string, password string) (*users.User, *errors.RestErroAPI) {

	request := users.LoginRequest{
		Email:    email,
		Password: password,
	}

	response := usersRestClient.Post("/users/login", request)

	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("invalid restclient response when trying to login user")
	}

	if response.StatusCode > 299 {
		var restErro errors.RestErroAPI

		err := json.Unmarshal(response.Bytes(), &restErro)

		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}

		return nil, &restErro
	}

	var user users.User

	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal users login response")
	}

	return &user, nil

}
