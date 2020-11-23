package main

import (
	"fmt"
	"github.com/gocql/gocql"
	_ "github.com/gocql/gocql"
)

func main(){
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "test_keyspace"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	defer session.Close()
	if err == nil {
		fmt.Println("Successfully connected to cassandra cluster.")
	} else {
		fmt.Println("Error connecting to the cluster")
	}

	Controller(session)

	//if err := session.Query('Insert int')
}
