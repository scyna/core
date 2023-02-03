package scyna

import (
	reflect "reflect"
	"time"

	"github.com/nats-io/nats.go"
	scyna_proto "github.com/scyna/core/proto/generated"
	"google.golang.org/protobuf/proto"
)

type EventHandler[R proto.Message] func(ctx *Context, data R)

type eventStream struct {
	sender    string
	receiver  string
	executors map[string]func(m *nats.Msg)
}

var eventStreams map[string]*eventStream = make(map[string]*eventStream)

func RegisterEvent[R proto.Message](sender string, channel string, handler EventHandler[R]) {
	if _, err := JetStream.StreamInfo(sender); err != nil {
		panic("No stream `" + sender + "`")
	}

	if _, err := JetStream.ConsumerInfo(sender, module); err != nil {
		panic("No consumer `" + module + "` for stream `" + sender + "`")
	}

	stream := createOrGetEventStream(sender)
	subject := sender + "." + channel
	var event R
	ref := reflect.New(reflect.TypeOf(event).Elem())
	event = ref.Interface().(R)

	stream.executors[subject] = func(m *nats.Msg) {
		var msg scyna_proto.Event
		if err := proto.Unmarshal(m.Data, &msg); err != nil {
			LOG.Error("Can not parse event data:" + err.Error())
			return
		}

		trace := Trace{
			ID:        ID.Next(),
			Type:      TRACE_EVENT,
			Path:      subject,
			SessionID: Session.ID(),
			Time:      time.Now(),
			ParentID:  msg.TraceID,
		}

		context := NewContext(trace.ID)

		if err := proto.Unmarshal(msg.Body, event); err == nil {
			handler(context, event)
		} else {
			LOG.Error("Error in parsing data:" + err.Error())
		}
		trace.Record()
	}
}

func (es *eventStream) start() {
	sub, err := JetStream.PullSubscribe("", es.receiver, nats.BindStream(es.sender))

	if err != nil {
		panic("Error in start event stream. Sender=" + es.sender + "- Receiver=" + es.receiver + " Error:" + err.Error())
	}

	go func() {
		for {
			if messages, err := sub.Fetch(1); err == nil {
				if len(messages) == 1 {
					m := messages[0]
					if executor, ok := es.executors[m.Subject]; ok {
						executor(m)
					}
					m.Ack()
				}
			}
		}
	}()
}

func createOrGetEventStream(sender string) *eventStream {
	if stream, ok := eventStreams[sender]; ok {
		return stream
	}

	stream := &eventStream{
		sender:    sender,
		receiver:  module,
		executors: make(map[string]func(m *nats.Msg)),
	}

	eventStreams[sender] = stream
	return stream
}

func startEventStreams() {
	for _, e := range eventStreams {
		e.start()
	}
}
