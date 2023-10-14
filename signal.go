package scyna

import (
	"log"

	"github.com/nats-io/nats.go"
	scyna_utils "github.com/scyna/core/utils"

	"google.golang.org/protobuf/proto"
)

type Scope int

const SCOPE_MODULE Scope = 1  // a instance of module
const SCOPE_SESSION Scope = 2 // all instances of module

type SignalHandler[R proto.Message] func(data R)

type signal_ struct {
	handler func(m *nats.Msg)
	subs    *nats.Subscription
	scope   Scope
}

var signals = make(map[string]signal_)

func init() { RegisterSetup(start) }

func start() {
	for channel, sig := range signals {
		var err error
		if sig.scope == SCOPE_MODULE {
			if sig.subs, err = Nats.QueueSubscribe(channel, Module(), sig.handler); err != nil {
				panic("Error in register Signal")
			}
		} else {
			if sig.subs, err = Nats.Subscribe(channel, sig.handler); err != nil {
				panic("Error in register Signal")
			}
		}

		if err != nil {
			panic("Error in register Signal")
		}
	}
}

func RegisterSignal[R proto.Message](channel string, handler SignalHandler[R], scope ...Scope) {
	log.Print("Register SignalLite:", channel)

	if len(scope) > 1 {
		panic("Invalid scope parametter")
	}

	signalScope := SCOPE_MODULE

	if len(scope) == 1 {
		signalScope = scope[0]
	}

	sig := scyna_utils.NewMessageForType[R]()

	cb := func(m *nats.Msg) {
		if err := proto.Unmarshal(m.Data, sig); err == nil {
			handler(sig)
		} else {
			Session.Error("Error in parsing data:" + err.Error())
		}
	}

	signals[channel] = signal_{
		handler: cb,
		scope:   signalScope,
	}
}

func EmitSignal(channel string, signal proto.Message) {
	if data, err := proto.Marshal(signal); err == nil {
		Nats.Publish(channel, data)
	}
}
