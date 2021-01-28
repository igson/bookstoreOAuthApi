package domain

import (
	"strings"
	"time"

	"github.com/igson/bookstoreOAuthApi/src/utils/errors"
)

const (
	expirationTime = 12
)

type Token struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() Token {
	return Token{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (t Token) IsExpired() bool {
	return time.Unix(t.Expires, 0).Before(time.Now().UTC())
}

func (t *Token) Validate() *errors.RestErroAPI {

	t.AccessToken = strings.TrimSpace(t.AccessToken)

	if t.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}
	if t.UserID <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if t.ClientID <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if t.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}
