package scyna

import (
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	scyna_proto "github.com/scyna/core/proto"
	scyna_utils "github.com/scyna/core/utils"
	"google.golang.org/protobuf/proto"
)

type EndpointHandler[R proto.Message] func(ctx *Endpoint, request R) Error

func RegisterEndpoint[R proto.Message](url string, handler EndpointHandler[R]) {
	log.Println("Register Endpoint: ", url)

	_, err := Nats.QueueSubscribe(scyna_utils.SubscriberURL(url), "API", func(m *nats.Msg) {
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
			if err := json.Unmarshal(ctx.Request.Body, request); err != nil {
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
	})

	if err != nil {
		panic("Can not register endpoint:" + url)
	}
}

func sendRequest(url string, request proto.Message, response proto.Message) Error {
	trace := CreateTrace(url, TRACE_ENDPOINT)
	return sendRequest_(trace, url, request, response)
}

func sendRequest_(trace *trace, url string, request proto.Message, response proto.Message) Error {
	defer trace.Record()

	req := scyna_proto.Request{TraceID: trace.ID, JSON: false}
	res := scyna_proto.Response{}

	if request != nil {
		var err error
		if req.Body, err = proto.Marshal(request); err != nil {
			return BAD_REQUEST
		}
	}

	if data, err := proto.Marshal(&req); err == nil {
		if msg, err := Nats.Request(scyna_utils.PublishURL(url), data, 10*time.Second); err == nil {
			if err := proto.Unmarshal(msg.Data, &res); err != nil {
				return SERVER_ERROR
			}
		} else {
			return SERVER_ERROR
		}
	} else {
		return BAD_REQUEST
	}

	trace.Status = uint32(res.Code)

	if res.Code == 200 {
		if err := proto.Unmarshal(res.Body, response); err == nil {
			return OK
		}
	} else {
		var ret scyna_proto.Error
		if err := proto.Unmarshal(res.Body, &ret); err == nil {
			return NewError(ret.Code, ret.Message)
		}
	}

	return SERVER_ERROR
}
