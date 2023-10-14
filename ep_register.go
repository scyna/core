package scyna

import (
	"log"

	"github.com/nats-io/nats.go"
	scyna_utils "github.com/scyna/core/utils"
	"google.golang.org/protobuf/proto"
)

type EndpointHandler[R proto.Message] func(ctx *Endpoint, request R) Error

type endpoint struct {
	handler func(m *nats.Msg)
	subs    *nats.Subscription
}

var endpoints = make(map[string]endpoint)

func init() { RegisterSetup(startEndpoints) }

func startEndpoints() {
	for url, ep := range endpoints {
		var err error
		if ep.subs, err = Nats.QueueSubscribe(scyna_utils.SubscriberURL(url), "API", ep.handler); err != nil {
			panic("Can not register endpoint:" + url)
		}
	}
}

func RegisterEndpoint[R proto.Message](url string, handler EndpointHandler[R]) {
	log.Println("Register endpoint:", url)
	if _, ok := endpoints[url]; ok {
		panic("Endpoint already registered: " + url)
	}

	f := func(m *nats.Msg) {
		request := scyna_utils.NewMessageForType[R]()
		ctx := Endpoint{
			Context: Context{ID: 0},
			flushed: false,
			Reply:   m.Reply,
		}

		if err := proto.Unmarshal(m.Data, &ctx.Request); err != nil {
			log.Print("Register unmarshal error response data:", err.Error())
			return
		}
		ctx.ID = ctx.Request.TraceID

		if ctx.Request.JSON {
			if err := proto.Unmarshal(ctx.Request.Body, request); err != nil {
				log.Print("Bad Request: " + err.Error())
				ctx.flushError(400, BAD_REQUEST)
			}
		} else {
			if err := proto.Unmarshal(ctx.Request.Body, request); err != nil {
				log.Print("Bad Request: " + err.Error())
				ctx.flushError(400, BAD_REQUEST)
			}
			ctx.request = request
		}

		e := handler(&ctx, request)
		if !ctx.flushed {
			if e == OK {
				ctx.flushError(200, OK)
			} else {
				ctx.flushError(400, e)
			}
		}
	}
	endpoints[url] = endpoint{handler: f}
}
