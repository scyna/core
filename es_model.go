package scyna

import (
	"google.golang.org/protobuf/proto"
)

type Model[T proto.Message] struct {
	ID      any
	Version int64
	Data    T
	Event   proto.Message
	store   *EventStore[T]
}

func (m *Model[T]) CommitAndProject(event proto.Message) Error {
	m.Event = event
	if err := m.store.updateWriteModel(m, event); err != nil {
		return err
	}

	m.store.updateReadModel(m.ID)
	return nil
}
