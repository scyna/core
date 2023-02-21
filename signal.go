package scyna

import (
	"log"

	"github.com/nats-io/nats.go"
	scyna_utils "github.com/scyna/core/utils"
	"google.golang.org/protobuf/proto"
)

type SignalHandler[R proto.Message] func(data R)

func RegisterSignal[R proto.Message](channel string, handler SignalHandler[R]) {
	log.Print("Register SignalLite:", channel)
	signal := scyna_utils.NewMessageForType[R]()

	if _, err := Connection.QueueSubscribe(channel, module, func(m *nats.Msg) {
		if err := proto.Unmarshal(m.Data, signal); err == nil {
			handler(signal)
		} else {
			Session.Error("Error in parsing data:" + err.Error())
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
