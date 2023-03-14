package scyna

import (
	"reflect"
)

type DomainEventHandler[E any] func(event E)

var domainEventQueue chan any = make(chan any)

type eventEntry struct {
	executors []func(event any)
}

var eventRegistrations map[reflect.Type]*eventEntry = make(map[reflect.Type]*eventEntry)

func RegisterDomainEvent[E any](handler DomainEventHandler[E]) {
	var tmp E
	myType := reflect.TypeOf(tmp)

	reg, ok := eventRegistrations[myType]
	if !ok {
		eventRegistrations[myType] = &eventEntry{}
	}

	reg.executors = append(reg.executors, func(event any) {
		val, _ := reflect.ValueOf(event).Interface().(E)
		handler(val)
	})
}

func startDomainEventLoop() {
	go func() {
		for event := range domainEventQueue {
			reg, ok := eventRegistrations[reflect.TypeOf(event)]
			if ok {
				for _, ex := range reg.executors {
					ex(event)
				}
			}
		}
	}()
}
