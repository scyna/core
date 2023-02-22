package scyna

import (
	"fmt"
	"log"
	"time"

	scyna_proto "github.com/scyna/core/proto/generated"
	"google.golang.org/protobuf/proto"
)

type Context interface {
	Logger
	PublishEvent(channel string, data proto.Message) Error
	ScheduleTask(channel string, start time.Time, interval int64, message proto.Message, loop uint64) (uint64, Error)
	SendRequest(url string, request proto.Message, response proto.Message) Error
	SaveTag(key string, value string)
	OK(r proto.Message) Error
	Response(r proto.Message)
	Authenticate(uid string, apps []string, r proto.Message)
}

type Endpoint struct {
	ID      uint64
	Request scyna_proto.Request
	Reply   string
	flushed bool
	request proto.Message
}

func NewContext(id uint64) *Endpoint {
	return &Endpoint{ID: id}
}

func (ctx *Endpoint) PublishEvent(channel string, data proto.Message) Error {
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

func (ctx *Endpoint) ScheduleTask(channel string, start time.Time, interval int64, message proto.Message, loop uint64) (uint64, Error) {

	task := scyna_proto.Task{TraceID: ctx.ID}
	if data, err := proto.Marshal(message); err != nil {
		return 0, BAD_DATA
	} else {
		task.Data = data
	}

	var response scyna_proto.StartTaskResponse
	if data, err := proto.Marshal(&task); err != nil {
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

func (ctx *Endpoint) SendRequest(url string, request proto.Message, response proto.Message) Error {
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

func (ctx *Endpoint) SaveTag(key string, value string) {
	if ctx.ID == 0 {
		return
	}
	EmitSignal(scyna_proto.TAG_CREATED_CHANNEL, &scyna_proto.TagCreatedSignal{
		TraceID: ctx.ID,
		Key:     key,
		Value:   value,
	})
}

func (l *Endpoint) writeLog(level LogLevel, message string) {
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

func (l *Endpoint) Info(messsage string) {
	l.writeLog(LOG_INFO, messsage)
}

func (l *Endpoint) Error(messsage string) {
	l.writeLog(LOG_ERROR, messsage)
}

func (l *Endpoint) Warning(messsage string) {
	l.writeLog(LOG_WARNING, messsage)
}

func (l *Endpoint) Debug(messsage string) {
	l.writeLog(LOG_DEBUG, messsage)
}

func (l *Endpoint) Fatal(messsage string) {
	l.writeLog(LOG_FATAL, messsage)
}
