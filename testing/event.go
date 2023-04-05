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
	input           R
}

func Event[R proto.Message](handler scyna.EventHandler[R]) *eventTest[R] {
	return &eventTest[R]{handler: handler}
}

func (t *eventTest[R]) WithEvent(event R) *eventTest[R] {
	t.input = event
	return t
}

func (t *eventTest[R]) ExpectEvent(event proto.Message, channel ...string) *eventTest[R] {
	if len(channel) > 1 {
		panic("Wrong parametter ")
	}
	if len(channel) == 1 {
		t.channel = channel[0]
	}

	t.event = event
	t.exactEventMatch = true
	return t
}

func (t *eventTest[R]) ExpectEventLike(event proto.Message, channel ...string) *eventTest[R] {
	if len(channel) > 1 {
		panic("Wrong parametter ")
	}
	if len(channel) == 1 {
		t.channel = channel[0]
	}

	t.event = event
	t.exactEventMatch = false
	return t
}

func (st *eventTest[R]) Run(t *testing.T) {
	streamName := scyna.Module()
	if len(st.channel) > 0 {
		createStream(streamName)
	}

	ctx := scyna.NewEvent(scyna.ID.Next())
	st.handler(ctx, st.input)

	if st.event != nil {
		if len(st.channel) > 0 {
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
		} else {
			time.Sleep(time.Millisecond * 100)
			receivedEvent := nextEvent()
			if receivedEvent == nil {
				t.Fatal("No event received")
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
		}
	}

	if len(st.channel) > 0 {
		deleteStream(streamName)
	}
}
