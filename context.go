package scyna

import (
	"fmt"
	"time"

	scyna_engine "github.com/scyna/core/engine"
	scyna_proto "github.com/scyna/core/proto/generated"
	"google.golang.org/protobuf/proto"
)

type Context struct {
	Logger
}

func NewContext(id uint64) *Context {
	return &Context{Logger{ID: id, session: false}}
}

func (ctx *Context) PublishEvent(channel string, data proto.Message) Error {
	msg := scyna_proto.Event{TraceID: ctx.ID}
	if data, err := proto.Marshal(data); err != nil {
		return BAD_DATA
	} else {
		msg.Body = data
	}

	if data, err := proto.Marshal(&msg); err != nil {
		return BAD_DATA
	} else {
		if _, err := JetStream.Publish(channel, data); err != nil {
			return STREAM_ERROR
		}
	}
	return nil
}

func (ctx *Context) ScheduleTask(task string, start time.Time, interval int64, data []byte, loop uint64) (Error, uint64) {
	var response scyna_proto.StartTaskResponse
	if err := ctx.SendRequest(scyna_engine.START_TASK_URL, &scyna_proto.StartTaskRequest{
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

func (ctx *Context) SaveTag(key string, value string) {
	if ctx.ID == 0 {
		return
	}
	EmitSignal(scyna_engine.TAG_CREATED_CHANNEL, &scyna_engine.TagCreatedSignal{
		TraceID: ctx.ID,
		Key:     key,
		Value:   value,
	})
}
