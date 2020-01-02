package token_acesso

import(
	"testing"	
	"time"
	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T)   {
	assert.EqualValues(t, 24, tempoExpiracao, "O token de acesso não pode ser vázio")
}
 
func TestGeraTokenDeAcesso(t *testing.T)   {
	token := GeraTokenDeAcesso()

	assert.False(t, token.IsTokenExpirado(), "O token de acesso não pode ser vázio")
	assert.EqualValues(t, "", token.AccessToken, "new access token should not have diferent")
	assert.True(t, token.UserId == 0, "access token should not have an associated user id")

}

func TestAcessoToken(t *testing.T)   {
	token := AccessToken{}
	
	assert.True(t, token.IsTokenExpirado(), "access token should not have an associated user id")
	
	token.ExpiredToken = time.Now().UTC().Add(tempoExpiracao * time.Hour).Unix()

}
