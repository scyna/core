package base

import (
	"fmt"
	"log"
	"strings"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type Connection struct {
	Nats   *nats.Conn
	Stream nats.JetStreamContext
}

func (c *Connection) SendRequest(url string, data proto.Message) *Error {
	/*TODO*/
	return nil
}

func (c *Connection) SendAndReceive(url string, request proto.Message, response proto.Message) *Error {
	/*TODO*/
	return nil
}

func (c *Connection) EmitSignal(channel string, signal proto.Message) error {
	data, err := proto.Marshal(signal)
	if err == nil {
		c.Nats.Publish(channel, data)
	}
	return err
}

func (c *Connection) Close() {
	c.Nats.Close()
}

func NewConnection(urls string, userName string, password string) *Connection {
	var err error
	var natsUrls []string
	for _, n := range strings.Split(urls, ",") {
		fmt.Printf("Nats configuration: nats://%s:4222\n", n)
		natsUrls = append(natsUrls, fmt.Sprintf("nats://%s:4222", n))
	}

	var conn *nats.Conn
	if userName != "" && password != "" {
		conn, err = nats.Connect(strings.Join(natsUrls, ","), nats.UserInfo(userName, password))
	} else {
		conn, err = nats.Connect(strings.Join(natsUrls, ","))
	}

	if err != nil {
		log.Fatal("Can not connect to NATS:", natsUrls)
	}

	/*init jetstream*/
	stream, err := conn.JetStream()
	if err != nil {
		panic("Init: " + err.Error())
	}

	return &Connection{Nats: conn, Stream: stream}
}
