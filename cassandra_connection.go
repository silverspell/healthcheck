package healthcheckmodule

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gocql/gocql"
)

type CassandraConnection struct {
	Addresses []string
	Keyspace  string
	User      string
	Pass      string
}

func (c *CassandraConnection) Connect() error {
	fmt.Println("Trying cassandra")
	cassandraHosts := os.Getenv("CASSANDRA_HOSTS")
	arrCassandraHosts := strings.Split(cassandraHosts, ",")
	cluster := gocql.NewCluster(arrCassandraHosts...)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: os.Getenv("CASSANDRA_USER"),
		Password: os.Getenv("CASSANDRA_PASS"),
	}
	cluster.Keyspace = os.Getenv("CASSANDRA_KEYSPACE")
	cluster.Consistency = gocql.Quorum
	cluster.ConnectTimeout = time.Duration(2 * time.Second)
	_, err := cluster.CreateSession()
	return err
}
