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
	subject := context + "." + channel
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

func (ctx *Context) Schedule(task string, start time.Time, interval int64, data []byte, loop uint64) (Error, uint64) {
	var response scyna_proto.StartTaskResponse
	if err := ctx.CallEndpoint(START_TASK_URL, &scyna_proto.StartTaskRequest{
		Context:  context,
		Topic:    fmt.Sprintf("%s.task.%s", context, task),
		Data:     data,
		Time:     start.Unix(),
		Interval: interval,
		Loop:     loop,
	}, &response); err != OK {
		return err, 0
	}

	return nil, response.Id
}

func (ctx *Context) CallEndpoint(url string, request proto.Message, response proto.Message) Error {
	trace := Trace{
		ID:       ID.Next(),
		ParentID: ctx.ID,
		Time:     time.Now(),
		Path:     url,
		Type:     TRACE_ENDPOINT,
		Source:   context,
	}
	return callEndpoint_(&trace, url, request, response)
}

func (ctx *Context) Tag(key string, value string) {
	if ctx.ID == 0 {
		return
	}
	emitSignal(TAG_CREATED_CHANNEL, &scyna_proto.TagCreatedSignal{
		TraceID: ctx.ID,
		Key:     key,
		Value:   value,
	})
}
