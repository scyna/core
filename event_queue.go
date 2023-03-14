package scyna

import (
	"log"
	"reflect"
	"time"

	"google.golang.org/protobuf/proto"
)

type DomainEventHandler[E any] func(ctx *Event, event E)

type eventItem struct {
	channel     string
	parentTrace uint64
	data        proto.Message
}

type eventRegistrationEntry struct {
	executors []func(event eventItem)
}

var eventQueue chan eventItem = make(chan eventItem)
var eventRegistrations map[string]*eventRegistrationEntry = make(map[string]*eventRegistrationEntry)

func RegisterDomainEvent[E proto.Message](channel string, handler DomainEventHandler[E]) {

	reg, ok := eventRegistrations[channel]
	if !ok {
		eventRegistrations[channel] = &eventRegistrationEntry{}
	}

	reg.executors = append(reg.executors, func(event eventItem) {
		val, ok := reflect.ValueOf(event.data).Interface().(E)
		if !ok {
			log.Print("Event type not match to EventHandler")
			return
		}

		trace := Trace{
			ID:        ID.Next(),
			Type:      TRACE_DOMAIN_EVENT,
			Path:      channel,
			SessionID: Session.ID(),
			Time:      time.Now(),
			ParentID:  event.parentTrace,
		}

		ctx := NewEvent(trace.ID)
		handler(ctx, val)

		trace.Record()
	})
}

func startDomainEventLoop() {
	go func() {
		for event := range eventQueue {
			reg, ok := eventRegistrations[event.channel]
			if ok {
				for _, executor := range reg.executors {
					executor(event)
				}
			}
		}
	}()
}
