package main

import (
	"log"
	"os"
	"time"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func init() {

	cluster := gocql.NewCluster("127.0.0.1", "..", "...")
	cluster.Port = 29042
	cluster.ProtoVersion = 4
	cluster.Consistency = gocql.Quorum

	cluster.Keyspace = "keyspacename"

	cluster.Authenticator = gocql.PasswordAuthenticator{Username: "username", Password: "password"}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	//you should have your certificate in your path
	cluster.SslOpts = &gocql.SslOptions{
		CaPath:                 dir + "/certificatename.pem",
		EnableHostVerification: true,
	}
	cluster.ConnectTimeout = time.Second * 20

	tempsession, err := cluster.CreateSession()
	if err != nil {
		log.Println("Error is happening in    cassandra file    ", err)
		return
	}
	Session = tempsession

	// //defer Session.Close()

}

func main() {

}
