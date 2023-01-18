package scyna

import (
	"fmt"

	"github.com/scylladb/gocqlx/v2/qb"
	"google.golang.org/protobuf/proto"
)

type eventStore struct {
	version     uint64
	storeQuery  string
	outboxQuery string
	publishing  bool
}

var EventStore *eventStore

func InitEventStore(keyspace string) {
	var version uint64 = 0

	if err := qb.Select(keyspace + ".event_store").
		Max("event_id").
		Query(DB).
		GetRelease(&version); err != nil {
		version = 1
		//log.Fatal("Can not init EventStore")
	}

	EventStore = &eventStore{
		version:     version,
		publishing:  false,
		storeQuery:  fmt.Sprintf("INSERT INTO %s.event_store(event_id, entity_id, channel, data) VALUES(?,?,?,?) ", keyspace),
		outboxQuery: fmt.Sprintf("INSERT INTO %s.outbox(event_id, trace_id) VALUES(?,?) ", keyspace),
	}
}

func (events *eventStore) Append(ctx *Command, aggregate uint64, channel string, event proto.Message) bool {
	if events == nil {
		panic("EventStore is not initialized")
	}

	var id = events.version + 1

	bytes, err := proto.Marshal(event)
	if err != nil {
		ctx.Logger.Error("Can not marshal event data")
		return false
	}

	ctx.Batch.Query(events.storeQuery, id, aggregate, channel, bytes)
	ctx.Batch.Query(events.outboxQuery, id, ctx.Request.TraceID)

	if err := DB.ExecuteBatch(ctx.Batch); err == nil {
		events.version = id
		events.publish()
		return true
	}

	return false
}

func (events *eventStore) publish() {
	if events.publishing {
		return
	}

	go func() {
		events.publishing = true

		/*TODO: publish*/

		events.publishing = false
	}()
}
