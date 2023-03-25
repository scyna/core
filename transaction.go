package scyna

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/qb"
	scyna_proto "github.com/scyna/core/proto/generated"
	"google.golang.org/protobuf/proto"
)

var _version uint64 = 0
var _query string

func InitSingleWriter(keyspace string) {
	var version uint64 = 1

	if err := qb.Select(keyspace + ".event_store").
		Max("event_id").
		Query(DB).
		GetRelease(&version); err != nil {
		panic("Error in load single writer configuration")
	}

	if version == 0 {
		version = 1
	}

	_version = version
	_query = fmt.Sprintf("INSERT INTO %s.event_store(event_id, entity_id, channel, data) VALUES(?,?,?,?) ", keyspace)
}

func NewTransaction(context Context) *transaction {
	return &transaction{
		batch:   DB.NewBatch(gocql.UnloggedBatch),
		entity:  0,
		context: context,
	}
}

type transaction struct {
	batch   *gocql.Batch
	channel string
	event   proto.Message
	entity  uint64
	context Context
}

func (trans *transaction) Batch() *gocql.Batch {
	return trans.batch
}

func (trans *transaction) SetEntity(id uint64) *transaction {
	trans.entity = id
	return trans
}

func (trans *transaction) SetEvent(event proto.Message) *transaction {
	trans.event = event
	return trans
}

func (trans *transaction) SetChannel(channel string) *transaction {
	trans.channel = channel
	return trans
}

func (trans *transaction) Commit() Error {
	if _version == 0 {
		panic("SingleWriter is not initialized")
	}

	if (trans.entity == 0) || (trans.event == nil) {
		return BAD_DATA
	}

	var id = _version + 1

	bytes, err := proto.Marshal(trans.event)
	if err != nil {
		trans.context.Error("Can not marshal event data")
		return BAD_DATA
	}

	trans.batch.Query(_query, id, trans.entity, trans.channel, bytes)

	if err := DB.ExecuteBatch(trans.batch); err != nil {
		trans.context.Error(err.Error())
		return SERVER_ERROR
	}

	_version = id

	if len(trans.channel) > 0 {

		eventMessage := &scyna_proto.Event{
			TraceID: trans.context.TraceID(),
			Entity:  trans.entity,
			Version: id,
		}

		if data, err := proto.Marshal(trans.event); err != nil {
			return BAD_DATA
		} else {
			eventMessage.Body = data
		}

		if data, err := proto.Marshal(eventMessage); err != nil {
			return BAD_DATA
		} else {
			if _, err := JetStream.Publish(buildSubject(module, trans.channel), data); err != nil {
				return STREAM_ERROR
			}
		}
	}
	return nil
}
