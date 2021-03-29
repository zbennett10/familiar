package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

func get() *gocql.Session {
	// Connect to Cassandra (hard-coded to development)
	cluster := gocql.NewCluster("localhost:9042")
	cluster.Keyspace = "familiar"

	session, err := cluster.CreateSession()
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("Connected to Cassandra\n")
	}

	return session
}
