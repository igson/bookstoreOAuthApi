package mysql_utils

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/igson/bookstoreUserApi/utils/erros"
)

const (
	errorNoRows    = "no rows in result set"
	indexUniqEmail = "email_UNIQUE"
)

func ParserError(erro error) *erros.MsgErroApi {

	driverErro, ok := erro.(*mysql.MySQLError)

	if !ok {
		if strings.Contains(erro.Error(), sql.ErrNoRows.Error()) {
			return erros.MsgNotFoundErro("Nenhum registro encontrado.")
		}
		fmt.Println("Erro: ", erro.Error())
		return erros.MsgInternalServerError("Erro de conexão.")
	}

	switch driverErro.Number {
	case 1062:
		return erros.MsgBadRequestErro("Dados inválidos")
	}
	return erros.MsgInternalServerError("Erro ao processar requisição")
}
