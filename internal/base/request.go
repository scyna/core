package base

import (
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type request struct {
	url     string
	request proto.Message
	conn    *nats.Conn
}

func (r *request) Receive(response proto.Message) *Error {
	/*TODO*/
	return nil
}

func (r *request) Send() *Error {
	/*TODO*/
	return nil
}
