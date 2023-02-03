package scyna_setup

import (
	"github.com/nats-io/nats.go"
	"github.com/scylladb/gocqlx/v2"
)

var Connection *nats.Conn
var JetStream nats.JetStreamContext
var DB gocqlx.Session

func init() {

}
