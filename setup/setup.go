package scyna_setup

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gocql/gocql"
	"github.com/nats-io/nats.go"
	"github.com/scylladb/gocqlx/v2"
)

var Connection *nats.Conn
var JetStream nats.JetStreamContext
var DB gocqlx.Session

func Init() {
	initNATS()
	initScylla([]string{"127.0.0.1"}, "", "", "")
}

func initNATS() {
	var err error
	var nats_ = []string{"nats://127.0.0.1:4222"}
	Connection, err = nats.Connect(strings.Join(nats_, ","))
	if err != nil {
		panic("Can not connect to NATS")
	}
	JetStream, err = Connection.JetStream()
	if err != nil {
		panic("Init: " + err.Error())
	}
}

func initScylla(host []string, username string, password string, location string) {
	cluster := gocql.NewCluster(host...)
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: username, Password: password}
	cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy(location)
	cluster.ConnectTimeout = time.Second * 10
	cluster.DisableInitialHostLookup = true
	cluster.Consistency = gocql.Quorum

	log.Printf("Connect to db: %s\n", host)

	var err error
	DB, err = gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		panic(fmt.Sprintf("Can not create session: Host = %s, Error = %s ", host, err.Error()))
	}
}
