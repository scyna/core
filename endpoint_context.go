package scyna

import (
	"encoding/json"
	"fmt"
	"log"

	scyna_proto "github.com/scyna/core/proto/generated"
	"google.golang.org/protobuf/proto"
)

type Endpoint struct {
	Context
	Request scyna_proto.Request
	Reply   string
	flushed bool
	request proto.Message
}

func NewEndpoint(id uint64) *Endpoint {
	return &Endpoint{Context: Context{ID: id}}
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

func (ctx *Endpoint) OK(r proto.Message) Error {
	ctx.Response(r)
	return OK
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

func (ctx *Endpoint) Authenticate(uid string, apps []string) (string, uint64, Error) {
	var auth scyna_proto.CreateAuthResponse
	if err := sendRequest(scyna_proto.AUTH_CREATE_URL,
		&scyna_proto.CreateAuthRequest{UID: uid, Apps: apps},
		&auth); err != OK {
		return "", 0, SERVER_ERROR
	}
	return auth.Token, auth.Expired, nil
}

func (ctx *Endpoint) AuthOK(token string, expired uint64, r proto.Message) Error {
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
	ctx.tag(uint32(response.Code), r)
	return OK
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
		Session.Error(fmt.Sprintf("Nats publish to [%s] error: %s", ctx.Reply, err.Error()))
	}
}

func (ctx *Endpoint) tag(code uint32, response proto.Message) {
	if ctx.ID == 0 {
		return
	}

	res, _ := json.Marshal(response)

	if ctx.Request.JSON {
		EmitSignal(scyna_proto.ENDPOINT_DONE_CHANNEL, &scyna_proto.EndpointDoneSignal{
			TraceID:  ctx.ID,
			Response: string(res),
			Request:  string(string(ctx.Request.Body)),
		})
	} else {
		req, _ := json.Marshal(ctx.request)
		EmitSignal(scyna_proto.ENDPOINT_DONE_CHANNEL, &scyna_proto.EndpointDoneSignal{
			TraceID:  ctx.ID,
			Response: string(res),
			Request:  string(req),
		})
	}
}
