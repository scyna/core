package eventstore

import (
	"github.com/scyna/core/internal/base"
	"google.golang.org/protobuf/proto"
)

type Projector[T proto.Message] func(T) *base.Error

type IProjection interface {
	Update(data []byte) *base.Error
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

func (p *Projection[T]) Update(data []byte) *base.Error {
	var model T
	if err := proto.Unmarshal(data, model); err != nil {
		return base.BAD_DATA
	}

	return p.Execute(model)
}

func (p *Projection[T]) ParseEvent(event []byte) proto.Message {
	ret := proto.Clone(p.eventType)
	if err := proto.Unmarshal(event, ret); err != nil {
		return nil
	}

	return ret
}
