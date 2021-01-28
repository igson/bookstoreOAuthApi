package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "O token de acesso não pode ser vázio")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "brand new access token should not be expired")
	assert.EqualValues(t, "", at.AccessToken, "new access token should not have defined access token id")
	assert.True(t, at.UserID == 0, "new access token should not have an associated user id")
}

func TestTokenIsExpired(t *testing.T) {
	token := Token{}
	assert.True(t, token.IsExpired(), "access token should not have an associated user id")
	token.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, token.IsExpired())
}
