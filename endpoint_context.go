package scyna

import (
	"encoding/json"
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

type Endpoint struct {
	Context
	Request Request
	Reply   string
	request proto.Message
}

func (ctx *Endpoint) Error(e *Error) {
	response := Response{Code: 400}

	var err error
	if ctx.Request.JSON {
		response.Body, err = json.Marshal(e)
	} else {
		response.Body, err = proto.Marshal(e)
	}

	if err != nil {
		response.Code = int32(500)
		response.Body = []byte(err.Error())
	}
	ctx.flush(&response)
	ctx.tag(uint32(response.Code), e)
}

func (ctx *Endpoint) Done(r proto.Message) {
	response := Response{Code: 200}

	var err error
	if ctx.Request.JSON {
		response.Body, err = json.Marshal(r)
	} else {
		response.Body, err = proto.Marshal(r)
	}
	if err != nil {
		response.Code = int32(500)
		response.Body = []byte(err.Error())
	}

	ctx.flush(&response)
	ctx.tag(uint32(response.Code), r)
}

func (ctx *Endpoint) AuthDone(r proto.Message, token string, expired uint64) {
	response := Response{Code: 200, Token: token, Expired: expired}

	var err error
	if ctx.Request.JSON {
		response.Body, err = json.Marshal(r)
	} else {
		response.Body, err = proto.Marshal(r)
	}
	if err != nil {
		response.Code = int32(500)
		response.Body = []byte(err.Error())
	}

	ctx.flush(&response)
	ctx.tag(200, r)
}

func (ctx *Endpoint) flush(response *Response) {
	response.SessionID = Session.ID()
	bytes, err := proto.Marshal(response)
	if err != nil {
		log.Print("Register marshal error response data:", err.Error())
		return
	}
	err = Connection.Publish(ctx.Reply, bytes)
	if err != nil {
		LOG.Error(fmt.Sprintf("Nats publish to [%s] error: %s", ctx.Reply, err.Error()))
	}
}

func (ctx *Endpoint) tag(code uint32, response proto.Message) {
	if ctx.ID == 0 {
		return
	}
	res, _ := json.Marshal(response)

	EmitSignal(ENDPOINT_DONE_CHANNEL, &EnpointDoneSignal{
		TraceID:  ctx.ID,
		Response: string(res),
	})
}
