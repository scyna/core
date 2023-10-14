package scyna

import (
	"reflect"

	scyna_utils "github.com/scyna/core/utils"
	"google.golang.org/protobuf/proto"
)

type Projector[M proto.Message, E proto.Message] func(M, E)

type IProjection interface {
	Update([]byte, []byte)
	ParseEvent([]byte) proto.Message
	EventName() string
}

type Projection[M proto.Message, E proto.Message] struct {
	Execute Projector[M, E]
}

func RegisterProjection[M proto.Message, E proto.Message](store *EventStore[M], projector Projector[M, E]) {
	projection := &Projection[M, E]{Execute: projector}
	store.registerProjection(projection)
}

func (p *Projection[M, E]) Update(data []byte, event []byte) {
	model := scyna_utils.NewMessageForType[M]()
	if err := proto.Unmarshal(data, model); err != nil {
		return
	}

	ev := scyna_utils.NewMessageForType[E]()
	if err := proto.Unmarshal(event, ev); err != nil {
		return
	}

	p.Execute(model, ev)
}

func (p *Projection[M, E]) ParseEvent(event []byte) proto.Message {
	var ret E
	if err := proto.Unmarshal(event, ret); err != nil {
		return nil
	}

	return ret
}

func (p *Projection[M, E]) EventName() string {
	var ev E
	return reflect.TypeOf(ev).Elem().Name()
}
