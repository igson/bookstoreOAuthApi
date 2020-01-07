package cassandra

import (
	"github.com/gocql/gocql"
	"fmt"
)

var (
	cluster *gocql.ClusterConfig
)

func init() {
	
	cluster = gocql.NewCluster("127.0.0.1")
	
	cluster.Keyspace = "oauth"
	cluster.Consistency=gocql.Quorum
	
	fmt.Println("Conexão realizada...")

}

func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}