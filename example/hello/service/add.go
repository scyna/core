package service

import (
	"ex/hello/proto"

	scyna "github.com/scyna/core"
)

func AddHandler(ctx *scyna.Endpoint, request *proto.AddRequest) scyna.Error {
	ctx.Logger.Info("Receive AddRequest")

	sum := request.A + request.B
	if sum > 100 {
		return ADD_RESULT_TOO_BIG
	}

	ctx.Response(&proto.AddResponse{Sum: sum})
	return scyna.OK
}
