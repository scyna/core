package scyna

import (
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type EndpointHandler[R proto.Message] func(ctx *Endpoint, request R)
type EndpointLiteHandler func(ctx *Endpoint)

func callEndpoint(url string, request proto.Message, response proto.Message) *Error {
	trace := Trace{
		ID:       ID.Next(),
		ParentID: 0,
		Time:     time.Now(),
		Path:     url,
		Type:     TRACE_ENDPOINT,
	}
	return callEndpoint_(&trace, url, request, response)
}

func RegisterEndpoint[R proto.Message](url string, handler EndpointHandler[R]) {
	log.Println("Register Endpoint: ", url)
	var request R

	ctx := Endpoint{
		Context: Context{Logger{session: false}},
		request: request,
	}

	_, err := Connection.QueueSubscribe(SubscriberURL(url), "API", func(m *nats.Msg) {
		if err := proto.Unmarshal(m.Data, &ctx.Request); err != nil {
			log.Print("Register unmarshal error response data:", err.Error())
			return
		}

		ctx.ID = ctx.Request.TraceID
		ctx.Reply = m.Reply
		ctx.Reset(ctx.ID)
		ref := request.ProtoReflect().New()
		request = ref.Interface().(R)

		if ctx.Request.JSON {
			if err := json.Unmarshal(ctx.Request.Body, request); err != nil {
				log.Print("Bad Request: " + err.Error())
				ctx.Error(BAD_REQUEST)
			} else {
				handler(&ctx, request)
			}

		} else {
			if err := proto.Unmarshal(ctx.Request.Body, request); err != nil {
				log.Print("Bad Request: " + err.Error())
				ctx.Error(BAD_REQUEST)
			} else {
				handler(&ctx, request)
			}
		}
	})

	if err != nil {
		Fatal("Can not register endpoint:", url)
	}
}

func RegisterEndpointLite(url string, handler EndpointLiteHandler) {
	log.Println("Register EndpointLite:", url)
	ctx := Endpoint{
		Context: Context{Logger{session: false}},
	}

	_, err := Connection.QueueSubscribe(SubscriberURL(url), "API", func(m *nats.Msg) {
		if err := proto.Unmarshal(m.Data, &ctx.Request); err != nil {
			log.Print("Register unmarshal error response data:", err.Error())
			return
		}

		ctx.ID = ctx.Request.TraceID
		ctx.Reply = m.Reply
		ctx.Reset(ctx.ID)
		handler(&ctx)
	})

	if err != nil {
		Fatal("Can not register command:", url)
	}
}

func callEndpoint_(trace *Trace, url string, request proto.Message, response proto.Message) *Error {
	defer trace.Record()

	req := Request{TraceID: trace.ID, JSON: false}
	res := Response{}

	if request != nil {
		var err error
		if req.Body, err = proto.Marshal(request); err != nil {
			return BAD_REQUEST
		}
	}

	if data, err := proto.Marshal(&req); err == nil {
		if msg, err := Connection.Request(PublishURL(url), data, 10*time.Second); err == nil {
			if err := proto.Unmarshal(msg.Data, &res); err != nil {
				return SERVER_ERROR
			}
		} else {
			return SERVER_ERROR
		}
	} else {
		return BAD_REQUEST
	}

	trace.SessionID = res.SessionID
	trace.Status = res.Code
	if res.Code == 200 {
		if err := proto.Unmarshal(res.Body, response); err == nil {
			return OK
		}
	} else {
		var ret Error
		if err := proto.Unmarshal(res.Body, &ret); err == nil {
			return &ret
		}
	}
	return SERVER_ERROR
}
