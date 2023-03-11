package scyna

import (
	"fmt"
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

type TaskBuilder struct {
	ctx     *context
	channel string
	message proto.Message
	time    time.Time
}

func (t *TaskBuilder) Channel(channel string) *TaskBuilder {
	t.channel = channel
	return t
}

func (t *TaskBuilder) Data(data proto.Message) *TaskBuilder {
	t.message = data
	return t
}

func (t *TaskBuilder) Time(time time.Time) *TaskBuilder {
	t.time = time
	return t
}

func (t *TaskBuilder) ScheduleOne() {
	t.schedule(1, 1000)
}

func (t *TaskBuilder) ScheduleSome(count uint64, interval int64) {
	t.schedule(count, interval)
}

func (t *TaskBuilder) ScheduleRepeat(interval int64) {
	t.schedule(0, interval)
}

func (t *TaskBuilder) schedule(count uint64, interval int64) (uint64, Error) {
	task := scyna_proto.Task{TraceID: t.ctx.ID}
	if data, err := proto.Marshal(t.message); err != nil {
		return 0, BAD_DATA
	} else {
		task.Data = data
	}

	var response scyna_proto.StartTaskResponse
	if data, err := proto.Marshal(&task); err != nil {
		return 0, BAD_DATA
	} else {
		if err := t.ctx.SendRequest(scyna_proto.START_TASK_URL, &scyna_proto.StartTaskRequest{
			Module:   module,
			Topic:    fmt.Sprintf("%s.%s", module, t.channel),
			Data:     data,
			Time:     t.time.Unix(),
			Interval: interval,
			Loop:     count,
		}, &response); err != OK {
			return 0, err
		}
	}
	return response.Id, nil

}
