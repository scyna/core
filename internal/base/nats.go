package base

import (
	"fmt"
	"log"
	"strings"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type Nats struct {
	Conn   *nats.Conn
	Stream nats.JetStreamContext
}

func (c *Nats) EmitSignal(channel string, signal proto.Message) error {
	data, err := proto.Marshal(signal)
	if err == nil {
		c.Conn.Publish(channel, data)
	}
	return err
}

func (c *Nats) Close() {
	c.Conn.Close()
}

func NewConnection(urls string, userName string, password string) *Nats {
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

	return &Nats{Conn: conn, Stream: stream}
}
