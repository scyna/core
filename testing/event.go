package scyna_test

import (
	"log"
	"testing"
	"time"

	scyna "github.com/scyna/core"
	scyna_proto "github.com/scyna/core/proto/generated"
	"google.golang.org/protobuf/proto"
)

type eventTest[R proto.Message] struct {
	channel         string
	event           proto.Message
	exactEventMatch bool
	handler         scyna.EventHandler[R]
}

func EventTest[R proto.Message](handler scyna.EventHandler[R]) *eventTest[R] {
	return &eventTest[R]{handler: handler}
}

func (t *eventTest[R]) OutputChannel(channel string) *eventTest[R] {
	t.channel = channel
	return t
}

func (t *eventTest[R]) ExpectOutputEvent(event R) *eventTest[R] {
	t.event = event
	t.exactEventMatch = true
	return t
}

func (t *eventTest[R]) MatchOutputEvent(event proto.Message) *eventTest[R] {
	t.event = event
	t.exactEventMatch = false
	return t
}

func (st *eventTest[R]) Run(t *testing.T, input R) {
	streamName := scyna.Module()
	if len(st.channel) > 0 {
		createStream(streamName)
	}

	ctx := scyna.NewEvent(scyna.ID.Next())
	st.handler(ctx, input)

	if st.event != nil {
		subs, err := scyna.JetStream.SubscribeSync(streamName + "." + st.channel)
		if err != nil {
			t.Fatal("Error in subscribe")
		}

		msg, err := subs.NextMsg(time.Second)
		if err != nil {
			t.Fatal("Timeout")
		}

		var event scyna_proto.Event
		if err := proto.Unmarshal(msg.Data, &event); err != nil {
			log.Print("Register unmarshal error response data:", err.Error())
			t.Fatal("Can not parse received event")
		}

		receivedEvent := proto.Clone(st.event)
		if proto.Unmarshal(event.Body, receivedEvent) != nil {
			t.Fatal("Can not parse received event")
		}

		if st.exactEventMatch {
			if !proto.Equal(st.event, receivedEvent) {
				t.Fatal("Event not match")
			}
		} else {
			if !matchMessage(st.event, receivedEvent) {
				t.Fatal("Event not match")
			}
		}

		subs.Unsubscribe()
	}

	if len(st.channel) > 0 {
		deleteStream(streamName)
	}
}
