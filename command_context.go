package scyna

import (
	"github.com/gocql/gocql"
	"google.golang.org/protobuf/proto"
)

type Command struct {
	Endpoint
	Batch *gocql.Batch
}

func (ctx *Command) StoreEvent(aggregate uint64, channel string, event proto.Message) Error {

	if !EventStore.Add(ctx, aggregate, channel, event) {
		return SERVER_ERROR
	}
	ctx.PostEvent(channel, event)
	return nil
}
