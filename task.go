package scyna

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
	scyna_proto "github.com/scyna/core/proto/generated"
	scyna_utils "github.com/scyna/core/utils"
	"google.golang.org/protobuf/proto"
)

type TaskHandler[R proto.Message] func(ctx Context, data R)

func RegisterTask[R proto.Message](sender string, channel string, handler TaskHandler[R]) {
	assureStreamReady(sender, module)
	stream := createOrGetEventStream(sender)
	subject := buildSubject(sender, channel)
	task := scyna_utils.NewMessageForType[R]()

	log.Print("Register task handler: ", subject)

	stream.executors[subject] = func(m *nats.Msg) {
		var msg scyna_proto.Task
		if err := proto.Unmarshal(m.Data, &msg); err != nil {
			Session.Error("Can not parse task data:" + err.Error())
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

		if err := proto.Unmarshal(msg.Data, task); err == nil {
			handler(context, task)
		} else {
			Session.Error("Error in parsing data:" + err.Error())
		}
		trace.Record()
	}
}
