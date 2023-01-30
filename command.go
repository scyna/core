package scyna

import (
	"encoding/json"
	"fmt"
	"log"
	reflect "reflect"

	"github.com/gocql/gocql"
	"github.com/nats-io/nats.go"
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
	Endpoint
	Batch *gocql.Batch
}

type CommandHandler[R proto.Message] func(ctx *Command, request R) Error

func RegisterCommand[R proto.Message](url string, handler CommandHandler[R]) {
	log.Println("Register Command: ", url)

	_, err := Connection.QueueSubscribe(SubscriberURL(url), "API", func(m *nats.Msg) {
		var request R
		ref := reflect.New(reflect.TypeOf(request).Elem())
		request = ref.Interface().(R)

		ctx := Command{
			Endpoint: Endpoint{
				Context: Context{Logger{session: false}},
				Reply:   m.Reply,
				flushed: false,
			},
			Batch: DB.NewBatch(gocql.UnloggedBatch),
		}

		if err := proto.Unmarshal(m.Data, &ctx.Request); err != nil {
			log.Print("Register unmarshal error response data:", err.Error())
			return
		}

		ctx.ID = ctx.Request.TraceID

		if ctx.Request.JSON {
			if err := json.Unmarshal(ctx.Request.Body, request); err != nil {
				log.Print("Bad Request: " + err.Error())
				ctx.flushError(400, BAD_REQUEST)
			}
		} else {
			if err := proto.Unmarshal(ctx.Request.Body, request); err != nil {
				log.Print("Bad Request: " + err.Error())
				ctx.flushError(400, BAD_REQUEST)
			}
		}

		e := handler(&ctx, request)
		if !ctx.flushed {
			if e == OK {
				ctx.flushError(200, OK)
			} else {
				ctx.flushError(400, e)
			}
		}
	})

	if err != nil {
		log.Fatal("Can not register command:", url)
	}
}

func (ctx *Command) StoreEvent(aggregate uint64, channel string, event proto.Message) Error {
	if _version == 0 {
		panic("SingleWriter is not initialized")
	}

	var id = _version + 1

	bytes, err := proto.Marshal(event)
	if err != nil {
		ctx.Logger.Error("Can not marshal event data")
		return BAD_DATA
	}

	ctx.Batch.Query(_query, id, aggregate, channel, bytes)

	if err := DB.ExecuteBatch(ctx.Batch); err != nil {
		ctx.Logger.Error(err.Error())
		return SERVER_ERROR
	}

	_version = id

	// if err := ctx.PublishEvent(channel, event); err != nil {
	// 	ctx.Logger.Error(err.Message())
	// 	/*TODO: system alert and panic here*/
	// 	return err
	// }
	return nil
}
