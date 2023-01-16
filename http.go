package scyna

import (
	"sync"

	scyna_proto "github.com/scyna/core/proto/generated"
)

type HttpContext struct {
	Request  scyna_proto.Request
	Response scyna_proto.Response
}

type HttpContextPool struct {
	sync.Pool
}

func (ctx *HttpContext) reset() {
	ctx.Request.Body = ctx.Request.Body[0:0]
	ctx.Request.TraceID = uint64(0)
	ctx.Response.Body = ctx.Response.Body[0:0]
	ctx.Response.Code = int32(0)
	ctx.Response.SessionID = uint64(0)
}

func newHttpContext() *HttpContext {
	return &HttpContext{
		Request: scyna_proto.Request{
			Body:    make([]byte, 4096),
			TraceID: 0,
		},
		Response: scyna_proto.Response{
			Body:      make([]byte, 0),
			SessionID: 0,
			Code:      200,
		},
	}
}

func (p *HttpContextPool) GetContext() *HttpContext {
	service, _ := p.Get().(*HttpContext)
	return service
}

func (p *HttpContextPool) PutContext(service *HttpContext) {
	service.reset()
	p.Put(service)
}

func NewContextPool() HttpContextPool {
	return HttpContextPool{
		sync.Pool{
			New: func() interface{} { return newHttpContext() },
		}}
}
