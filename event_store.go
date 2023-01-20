package scyna

import (
	"fmt"

	"github.com/scylladb/gocqlx/v2/qb"
	"google.golang.org/protobuf/proto"
)

type eventStore struct {
	version    uint64
	storeQuery string
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
		version:    version,
		storeQuery: fmt.Sprintf("INSERT INTO %s.event_store(event_id, entity_id, channel, data) VALUES(?,?,?,?) ", keyspace),
	}
}

func (events *eventStore) Append(ctx *Command, aggregate uint64, channel string, event proto.Message) Error {
	if events == nil {
		panic("EventStore is not initialized")
	}

	var id = events.version + 1

	bytes, err := proto.Marshal(event)
	if err != nil {
		ctx.Logger.Error("Can not marshal event data")
		return BAD_DATA
	}

	ctx.Batch.Query(events.storeQuery, id, aggregate, channel, bytes)

	if err := DB.ExecuteBatch(ctx.Batch); err != nil {
		ctx.Logger.Error(err.Error())
		return SERVER_ERROR
	}

	events.version = id

	// if err := ctx.PublishEvent(channel, event); err != nil {
	// 	ctx.Logger.Error(err.Message())
	// 	/*TODO: system alert and panic here*/
	// 	return err
	// }

	return nil
}
