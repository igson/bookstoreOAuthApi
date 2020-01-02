package token_acesso

import(
	"time"
	"testing"	
)

func TestGeraTokenDeAcesso(t *testing.T)  
	token := GeraTokenDeAcesso()
	if token.IsToken() {
		t.Error("O token de acesso não pode ser vázio")
	}

	if token.AcAccessToken != "" {
		t.Error("new accesse token should not have defined access token id")
	}

	if token.UserId != 0 {
		t.Error("new accesse token should not have an associated user id")
	}

}
