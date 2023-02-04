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
		if _, err := JetStream.Publish(buildSubject(module, channel), data); err != nil {
			return STREAM_ERROR
		}
	}
	return nil
}

func (ctx *Context) ScheduleTask(channel string, start time.Time, interval int64, message proto.Message, loop uint64) (uint64, Error) {
	var response scyna_proto.StartTaskResponse
	if data, err := proto.Marshal(message); err != nil {
		return 0, BAD_DATA
	} else {
		if err := ctx.SendRequest(scyna_proto.START_TASK_URL, &scyna_proto.StartTaskRequest{
			Module:   module,
			Topic:    fmt.Sprintf("%s.%s", module, channel),
			Data:     data,
			Time:     start.Unix(),
			Interval: interval,
			Loop:     loop,
		}, &response); err != OK {
			return 0, err
		}
	}
	return response.Id, nil
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
	EmitSignal(scyna_proto.TAG_CREATED_CHANNEL, &scyna_proto.TagCreatedSignal{
		TraceID: ctx.ID,
		Key:     key,
		Value:   value,
	})
}
