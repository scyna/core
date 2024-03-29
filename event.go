package scyna

import (
	"log"

	"github.com/nats-io/nats.go"
	scyna_proto "github.com/scyna/core/proto"
	scyna_utils "github.com/scyna/core/utils"
	"google.golang.org/protobuf/proto"
)

type Event struct {
	Context
}

func NewEvent(id uint64) *Event {
	return &Event{Context: Context{ID: id}}
}

type EventHandler[R proto.Message] func(ctx *Event, data R)

type eventStream struct {
	sender    string
	receiver  string
	executors map[string]func(m *nats.Msg)
}

var eventStreams map[string]*eventStream = make(map[string]*eventStream)

func RegisterEvent[R proto.Message](sender string, channel string, handler EventHandler[R]) {
	stream := createOrGetEventStream(sender)
	subject := buildSubject(sender, channel)
	event := scyna_utils.NewMessageForType[R]()

	log.Print("Register event handler: ", subject)

	stream.executors[subject] = func(m *nats.Msg) {
		var msg scyna_proto.Event
		if err := proto.Unmarshal(m.Data, &msg); err != nil {
			Session.Error("Can not parse event data:" + err.Error())
			return
		}
		trace := createTrace(subject, TRACE_EVENT, msg.TraceID)

		context := &Event{Context: Context{ID: trace.ID}}

		if err := proto.Unmarshal(msg.Body, event); err == nil {
			handler(context, event)
		} else {
			Session.Error("Error in parsing data:" + err.Error())
		}
		trace.record()
	}
}

func (es *eventStream) start() {
	assureStreamReady(es.sender, Module())
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

func init() { RegisterSetup(startEventStreams) }

func startEventStreams() {
	if testMode {
		return
	}

	for _, e := range eventStreams {
		e.start()
	}
}

func assureStreamReady(sender, receiver string) {
	if _, err := JetStream.StreamInfo(sender); err != nil {
		panic("No stream `" + sender + "`")
	}

	if _, err := JetStream.ConsumerInfo(sender, receiver); err != nil {
		panic("No consumer `" + module + "` for stream `" + sender + "`")
	}
}

func buildSubject(sender, channel string) string {
	return sender + "." + channel
}
