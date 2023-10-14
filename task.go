package scyna

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	scyna_const "github.com/scyna/core/const"
	scyna_proto "github.com/scyna/core/proto"
	scyna_utils "github.com/scyna/core/utils"
	"google.golang.org/protobuf/proto"
)

type TaskHandler[R proto.Message] func(ctx *Context, data R)

func RegisterTask[R proto.Message](sender string, channel string, handler TaskHandler[R]) {
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
		trace := createTrace(subject, TRACE_TASK, msg.TraceID)

		context := &Context{ID: trace.ID}

		if err := proto.Unmarshal(msg.Data, task); err == nil {
			handler(context, task)
		} else {
			Session.Error("Error in parsing data:" + err.Error())
		}
		trace.record()
	}
}

type taskBuilder struct {
	ctx     *Context
	channel string
	message proto.Message
	time    time.Time
}

func (t *taskBuilder) Data(data proto.Message) *taskBuilder {
	t.message = data
	return t
}

func (t *taskBuilder) Time(time time.Time) *taskBuilder {
	t.time = time
	return t
}

func (t *taskBuilder) ScheduleOne() Error {
	return t.schedule(1, 1000)
}

func (t *taskBuilder) ScheduleSome(count uint64, interval int64) Error {
	return t.schedule(count, interval)
}

func (t *taskBuilder) ScheduleRepeat(interval int64) Error {
	return t.schedule(0, interval)
}

func (t *taskBuilder) schedule(count uint64, interval int64) Error {
	task := scyna_proto.Task{TraceID: t.ctx.ID}
	if data, err := proto.Marshal(t.message); err != nil {
		return BAD_DATA
	} else {
		task.Data = data
	}

	var response scyna_proto.StartTaskResponse
	if data, err := proto.Marshal(&task); err != nil {
		return BAD_DATA
	} else {
		if err := t.ctx.SendRequest(scyna_const.START_TASK_URL, &scyna_proto.StartTaskRequest{
			Module:   module,
			Topic:    fmt.Sprintf("%s.%s", module, t.channel),
			Data:     data,
			Time:     t.time.Unix(),
			Interval: interval,
			Loop:     count,
		}, &response); err != OK {
			return err
		}
	}
	return nil
}
