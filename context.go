package scyna

import (
	"log"

	scyna_proto "github.com/scyna/core/proto"
	"google.golang.org/protobuf/proto"
)

type Context struct {
	ID uint64
}

func (ctx *Context) TraceID() uint64 {
	return ctx.ID
}

func (ctx *Context) Task(channel string) *taskBuilder {
	return &taskBuilder{ctx: ctx, channel: channel}
}

func (ctx *Context) RaiseEvent(event proto.Message) {
	go func() {
		eventQueue <- eventItem{Data: event, parentTrace: ctx.ID}
	}()
}

func (ctx *Context) PublishEvent(channel string, data proto.Message) Error {
	event := scyna_proto.Event{TraceID: ctx.ID}
	if data, err := proto.Marshal(data); err != nil {
		return BAD_DATA
	} else {
		event.Body = data
	}

	if data, err := proto.Marshal(&event); err != nil {
		return BAD_DATA
	} else {
		if _, err := JetStream.Publish(buildSubject(module, channel), data); err != nil {
			return STREAM_ERROR
		}
	}
	return nil
}

func (ctx *Context) SendRequest(url string, request proto.Message, response proto.Message) Error {
	trace := CreateTrace(url, TRACE_ENDPOINT, ctx.ID)
	return sendRequest_(trace, url, request, response)
}

func (l *Context) writeLog(level LogLevel, message string) {
	message = formatLog(message)
	log.Print(message)
	if l.ID > 0 {
		AddLog(LogData{
			ID:       l.ID,
			Sequence: Session.NextSequence(),
			Level:    level,
			Message:  message,
		})
	}
}

func (l *Context) Info(messsage string) {
	l.writeLog(LOG_INFO, messsage)
}

func (l *Context) Error(messsage string) {
	l.writeLog(LOG_ERROR, messsage)
}

func (l *Context) Warning(messsage string) {
	l.writeLog(LOG_WARNING, messsage)
}

func (l *Context) Debug(messsage string) {
	l.writeLog(LOG_DEBUG, messsage)
}

func (l *Context) Fatal(messsage string) {
	l.writeLog(LOG_FATAL, messsage)
}
