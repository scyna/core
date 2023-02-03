package scyna

import (
	"log"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type SignalHandler[R proto.Message] func(data R)

func RegisterSignal[R proto.Message](channel string, handler SignalHandler[R]) {
	log.Print("Register SignalLite:", channel)
	signal := newMessageForType[R]()

	if _, err := Connection.QueueSubscribe(channel, module, func(m *nats.Msg) {
		if err := proto.Unmarshal(m.Data, signal); err == nil {
			handler(signal)
		} else {
			LOG.Error("Error in parsing data:" + err.Error())
		}
	}); err != nil {
		panic("Error in register SignalLite")
	}
}

func EmitSignal(channel string, signal proto.Message) {
	if data, err := proto.Marshal(signal); err == nil {
		Connection.Publish(channel, data)
	}
}
