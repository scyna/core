package scyna

import (
	"fmt"
	"time"

	scyna_proto "github.com/scyna/core/proto/generated"
	"google.golang.org/protobuf/proto"
)

type Context struct {
	Logger
}

func (ctx *Context) PublishEvent(channel string, data proto.Message) Error {
	subject := module + "." + channel
	msg := scyna_proto.Event{TraceID: ctx.ID}
	if data, err := proto.Marshal(data); err != nil {
		return BAD_DATA
	} else {
		msg.Body = data
	}

	if data, err := proto.Marshal(&msg); err != nil {
		return BAD_DATA
	} else {
		if _, err := JetStream.Publish(subject, data); err != nil {
			return STREAM_ERROR
		}
	}
	return nil
}

func (ctx *Context) ScheduleTask(task string, start time.Time, interval int64, data []byte, loop uint64) (Error, uint64) {
	var response scyna_proto.StartTaskResponse
	if err := ctx.SendRequest(START_TASK_URL, &scyna_proto.StartTaskRequest{
		Module:   module,
		Topic:    fmt.Sprintf("%s.task.%s", module, task),
		Data:     data,
		Time:     start.Unix(),
		Interval: interval,
		Loop:     loop,
	}, &response); err != OK {
		return err, 0
	}

	return nil, response.Id
}

func (ctx *Context) SendRequest(url string, request proto.Message, response proto.Message) Error {
	trace := Trace{
		ID:       ID.Next(),
		ParentID: ctx.ID,
		Time:     time.Now(),
		Path:     url,
		Type:     TRACE_ENDPOINT,
		Source:   module,
	}
	return sendRequest_(&trace, url, request, response)
}

func (ctx *Context) Tag(key string, value string) {
	if ctx.ID == 0 {
		return
	}
	EmitSignal(TAG_CREATED_CHANNEL, &scyna_proto.TagCreatedSignal{
		TraceID: ctx.ID,
		Key:     key,
		Value:   value,
	})
}
