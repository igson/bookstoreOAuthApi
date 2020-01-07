package cassandra

import (
	"github.com/gocql/gocql"
	"fmt"
)

var (
	session *gocql.Session
)

func init() {
	
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency=gocql.Quorum
	
	var erro error

	if session, erro = cluster.CreateSession(); erro != nil {
			panic(erro)
	}

	fmt.Println("Conexão realizada...")

}

func GetSession() *gocql.Session {
	return session
}