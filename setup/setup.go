package scyna_setup

import (
	"strings"

	"github.com/nats-io/nats.go"
	"github.com/scyna/core/internal/base"
)

var Connection *nats.Conn
var JetStream nats.JetStreamContext
var DB *base.DB

func Init() {
	initNATS()
	DB = base.NewDB([]string{"127.0.0.1"}, "", "", "")
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
