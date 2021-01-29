package domain

import (
	"fmt"
	"strings"
	"time"

	"github.com/igson/bookstoreOAuthApi/src/utils/errors"
	"github.com/igson/bookstoreUsersApi/utils/crypto"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grandTypeClientCredentials = "client_credentials"
)

type Token struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	// Used for client_credentials grant type
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func GetNewAccessToken(userID int64) Token {
	return Token{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (t Token) IsExpired() bool {
	return time.Unix(t.Expires, 0).Before(time.Now().UTC())
}

func (at *Token) Generate() {
	at.AccessToken = crypto.GetMD5(fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires))
}

func (at *AccessTokenRequest) Validate() *errors.RestErroAPI {
	switch at.GrantType {
	case grantTypePassword:
		break

	case grandTypeClientCredentials:
		break

	default:
		return errors.NewBadRequestError("invalid grant_type parameter")
	}

	return nil
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
