package eventstore

import (
	"github.com/scyna/core/internal/base"
	"google.golang.org/protobuf/proto"
)

type Model[T proto.Message] struct {
	ID      any
	Version int64
	store   *EventStore[T]
	Data    T
	Event   proto.Message
}

func (m *Model[T]) CommitAndProject(event proto.Message) *base.Error {
	m.Event = event
	if err := m.store.UpdateWriteModel(m, event); err != nil {
		return err
	}
	/*TODO: update read model here*/
	return nil
}

// public void CommitAndProject(IMessage event_)
// {
// 	this.Event = event_;
// 	store.UpdateWriteModel(this, event_);
// 	store.UpdateReadModel(ID);
// }

// public void CommitAndPublish(string channel, IMessage event_)
// {
// 	this.Event = event_;
// 	store.UpdateWriteModel(this, event_);
// 	store.PublishToEventStream(channel, event_);
// }
