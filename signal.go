package scyna

import (
	"log"

	"github.com/nats-io/nats.go"
	scyna_utils "github.com/scyna/core/utils"
	"google.golang.org/protobuf/proto"
)

type SignalHandler[R proto.Message] func(data R)

type SignalScope int

const SIGNAL_SCOPE_MODULE SignalScope = 1  // a instance of module
const SIGNAL_SCOPE_SESSION SignalScope = 2 // all instances of module

func RegisterSignal[R proto.Message](channel string, handler SignalHandler[R], scope ...SignalScope) {
	log.Print("Register SignalLite:", channel)

	if len(scope) > 1 {
		panic("Invalid scope parametter")
	}
	signalScope := SIGNAL_SCOPE_MODULE

	if len(scope) == 1 {
		signalScope = scope[0]
	}

	signal := scyna_utils.NewMessageForType[R]()

	cb := func(m *nats.Msg) {
		if err := proto.Unmarshal(m.Data, signal); err == nil {
			handler(signal)
		} else {
			Session.Error("Error in parsing data:" + err.Error())
		}
	}

	if signalScope == SIGNAL_SCOPE_MODULE {
		if _, err := Nats.QueueSubscribe(channel, module, cb); err != nil {
			panic("Error in register SignalLite")
		}
	} else {
		if _, err := Nats.Subscribe(channel, cb); err != nil {
			panic("Error in register SignalLite")
		}
	}
}

func EmitSignal(channel string, signal proto.Message) {
	if data, err := proto.Marshal(signal); err == nil {
		Nats.Publish(channel, data)
	}
}
