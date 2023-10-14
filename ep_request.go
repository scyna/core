package scyna

import (
	"time"

	scyna_proto "github.com/scyna/core/proto"
	scyna_utils "github.com/scyna/core/utils"
	"google.golang.org/protobuf/proto"
)

func sendRequest(url string, request proto.Message, response proto.Message) Error {
	trace := createTrace(url, TRACE_ENDPOINT)
	return sendRequest_(trace, url, request, response)
}

func sendRequest_(trace *trace, url string, request proto.Message, response proto.Message) Error {
	defer trace.record()

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
