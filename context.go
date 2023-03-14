package scyna

import (
	"log"
	"time"

	scyna_proto "github.com/scyna/core/proto/generated"
	"google.golang.org/protobuf/proto"
)

type Context interface {
	Logger
	PublishEvent(channel string, data proto.Message) Error
	SendRequest(url string, request proto.Message, response proto.Message) Error
	Tag(key string, value string)
	OK(r proto.Message) Error
	Response(r proto.Message)
	Authenticate(uid string, apps []string, r proto.Message)
	TraceID() uint64
	Task(channel string) *TaskBuilder
}

type context struct {
	ID uint64
}

func NewEndpoint(id uint64) *Endpoint {
	return &Endpoint{context: context{ID: id}}
}

func NewEvent(id uint64) *Event {
	return &Event{context: context{ID: id}}
}

func (ctx *context) TraceID() uint64 {
	return ctx.ID
}

func (ctx *context) Task(channel string) *TaskBuilder {
	return &TaskBuilder{ctx: ctx, channel: channel}
}

func (ctx *context) RaiseDomainEvent(event any) {
	domainEventQueue <- event
}

func (ctx *context) PublishEvent(channel string, data proto.Message) Error {
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

func (ctx *context) SendRequest(url string, request proto.Message, response proto.Message) Error {
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

func (ctx *context) Tag(key string, value string) {
	if ctx.ID == 0 {
		return
	}
	EmitSignal(scyna_proto.TAG_CREATED_CHANNEL, &scyna_proto.TagCreatedSignal{
		TraceID: ctx.ID,
		Key:     key,
		Value:   value,
	})
}

func (l *context) writeLog(level LogLevel, message string) {
	message = formatLog(message)
	log.Print(message)
	if l.ID > 0 {
		AddLog(LogData{
			ID:       l.ID,
			Sequence: Session.NextSequence(),
			Level:    level,
			Message:  message,
			Session:  false,
		})
	}
}

func (l *context) Info(messsage string) {
	l.writeLog(LOG_INFO, messsage)
}

func (l *context) Error(messsage string) {
	l.writeLog(LOG_ERROR, messsage)
}

func (l *context) Warning(messsage string) {
	l.writeLog(LOG_WARNING, messsage)
}

func (l *context) Debug(messsage string) {
	l.writeLog(LOG_DEBUG, messsage)
}

func (l *context) Fatal(messsage string) {
	l.writeLog(LOG_FATAL, messsage)
}
