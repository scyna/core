package scyna

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type TaskHandler[R proto.Message] func(ctx *Context, data R)

func RegisterTask[R proto.Message](sender string, channel string, handler TaskHandler[R]) {
	subject := sender + ".task." + channel
	durable := "task_" + channel
	LOG.Info(fmt.Sprintf("Task: Channel %s, durable: %s", subject, durable))

	task := newMessageForType[R]()

	sub, err := JetStream.PullSubscribe(subject, durable, nats.BindStream(module))

	if err != nil {
		Fatal("Error in start event stream:", err.Error())
	}

	go func() {
		for {
			messages, err := sub.Fetch(1)
			if err != nil || len(messages) != 1 {
				continue
			}
			m := messages[0]

			trace := Trace{
				Path:      subject,
				SessionID: Session.ID(),
				Type:      TRACE_TASK,
				Time:      time.Now(),
				ID:        ID.Next(),
			}

			context := NewContext(trace.ID)

			if err := proto.Unmarshal(m.Data, task); err != nil {
				log.Print("Error in parsing data:", err)
				m.Ack()
				continue
			}

			handler(context, task)
			m.Ack()
			trace.Record()
		}
	}()
}
