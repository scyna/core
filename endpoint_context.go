package scyna

import (
	"encoding/json"
	"fmt"
	"log"

	scyna_engine "github.com/scyna/core/engine"
	scyna_proto "github.com/scyna/core/proto/generated"
	"google.golang.org/protobuf/proto"
)

type Endpoint struct {
	Context
	Request scyna_proto.Request
	Reply   string
	flushed bool
}

func (ctx *Endpoint) flushError(code int32, e Error) {
	response := scyna_proto.Response{Code: code}

	e_ := &scyna_proto.Error{
		Code:    e.Code(),
		Message: e.Message(),
	}

	var err error
	if ctx.Request.JSON {
		response.Body, err = json.Marshal(e_)
	} else {
		response.Body, err = proto.Marshal(e_)
	}

	if err != nil {
		response.Code = int32(500)
		response.Body = []byte(err.Error())
	}
	ctx.flush(&response)
	ctx.tag(uint32(response.Code), e_)
}

func (ctx *Endpoint) Response(r proto.Message) {
	response := scyna_proto.Response{Code: 200}

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
	response := scyna_proto.Response{Code: 200, Token: token, Expired: expired}

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

func (ctx *Endpoint) flush(response *scyna_proto.Response) {
	defer func() {
		ctx.flushed = true
	}()
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

	EmitSignal(scyna_engine.ENDPOINT_DONE_CHANNEL, &scyna_engine.EndpointDoneSignal{
		TraceID:  ctx.ID,
		Response: string(res),
	})
}
