package scyna_test

import (
	"log"

	scyna "github.com/scyna/core"
	"google.golang.org/protobuf/proto"
)

var events []proto.Message = make([]proto.Message, 0)

func Init(module string) {
	scyna.TestInit(scyna.RemoteConfig{
		ManagerUrl: "http://127.0.0.1:8081",
		Name:       module,
		Secret:     "123456",
	})

	log.Print(scyna.Session.ID())
	scyna.UseDirectLog(1)
	startEventLoop()
}

func Release() {
	scyna.Release()
}

func startEventLoop() {
	eventQueue := scyna.EventQueue()
	go func() {
		for event := range eventQueue {
			events = append(events, event.Data)
		}
	}()
}

func nextEvent() proto.Message {
	if len(events) == 0 {
		return nil
	}
	ret := events[0]
	events = events[1:]
	return ret
}
