package scyna

import (
	"log"
	"reflect"
	"time"

	"google.golang.org/protobuf/proto"
)

type DomainEventHandler[E any] func(ctx *Event, event E)

type eventItem struct {
	parentTrace uint64
	data        proto.Message
}

type eventRegistrationEntry struct {
	executors []func(event eventItem)
}

var eventQueue chan eventItem = make(chan eventItem)
var eventRegistrations map[reflect.Type]*eventRegistrationEntry = make(map[reflect.Type]*eventRegistrationEntry)

func RegisterDomainEvent[E proto.Message](handler DomainEventHandler[E]) {
	var tmp E
	t := reflect.TypeOf(tmp)

	reg, ok := eventRegistrations[t]
	if !ok {
		eventRegistrations[t] = &eventRegistrationEntry{}
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
			reg, ok := eventRegistrations[reflect.TypeOf(event.data)]
			if ok {
				for _, executor := range reg.executors {
					executor(event)
				}
			} else {
				log.Print("No handler attached to event")
			}
		}
	}()
}
