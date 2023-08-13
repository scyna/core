package eventstore

import (
	"log"

	"google.golang.org/protobuf/proto"
)

type Projector[T proto.Message] func(T) error

type IProjection interface {
	Update(data []byte)
	ParseEvent([]byte) proto.Message
}

type Projection[T proto.Message] struct {
	Execute   Projector[T]
	eventType proto.Message
}

func NewProjection[T proto.Message](execute Projector[T], eventType proto.Message) *Projection[T] {
	return &Projection[T]{
		Execute:   execute,
		eventType: eventType,
	}
}

func (p *Projection[T]) Update(data []byte) {
	var model T
	if err := proto.Unmarshal(data, model); err != nil {
		return
	}

	if err := p.Execute(model); err != nil {
		log.Print("projection/Update:", err)
		return
	}
}

func (p *Projection[T]) ParseEvent(event []byte) proto.Message {
	ret := proto.Clone(p.eventType)
	if err := proto.Unmarshal(event, ret); err != nil {
		return nil
	}

	return ret
}
