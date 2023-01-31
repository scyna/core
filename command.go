package scyna

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/qb"
	"google.golang.org/protobuf/proto"
)

var _version uint64 = 0
var _query string

func InitSingleWriter(keyspace string) {
	var version uint64 = 0

	if err := qb.Select(keyspace + ".event_store").
		Max("event_id").
		Query(DB).
		GetRelease(&version); err != nil {
		if err == gocql.ErrNotFound {
			version = 1
		} else {
			panic("Error in load single writer configuration")
		}
	}
	_version = version
	_query = fmt.Sprintf("INSERT INTO %s.event_store(event_id, entity_id, channel, data) VALUES(?,?,?,?) ", keyspace)
}

type Command struct {
	batch   *gocql.Batch
	channel string
	event   proto.Message
	entity  uint64
	context *Context
}

func (command *Command) Batch() *gocql.Batch {
	return command.batch
}

func NewCommand(context *Context) *Command {
	return &Command{
		batch:   DB.NewBatch(gocql.UnloggedBatch),
		entity:  0,
		context: context,
	}
}

func (command *Command) SetAggregateID(id uint64) *Command {
	command.entity = id
	return command
}

func (command *Command) SetEvent(event proto.Message) *Command {
	command.event = event
	return command
}

func (command *Command) SetChannel(channel string) *Command {
	command.channel = channel
	return command
}

func (command *Command) Commit() Error {
	if _version == 0 {
		panic("SingleWriter is not initialized")
	}

	if (command.entity == 0) || (command.event == nil) {
		return BAD_DATA
	}

	var id = _version + 1

	bytes, err := proto.Marshal(command.event)
	if err != nil {
		command.context.Error("Can not marshal event data")
		return BAD_DATA
	}

	command.batch.Query(_query, id, command.entity, command.channel, bytes)

	if err := DB.ExecuteBatch(command.batch); err != nil {
		command.context.Error(err.Error())
		return SERVER_ERROR
	}

	_version = id

	if len(command.channel) > 0 {
		if err := command.context.PublishEvent(command.channel, command.event); err != nil {
			command.context.Error(err.Message())
			/*TODO: system alert and panic here*/
			return err
		}
	}
	return nil
}
