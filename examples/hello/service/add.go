package hello

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/example/hello/proto"
)

func Add(s *scyna.Endpoint, request *proto.AddRequest) scyna.Error {
	s.Logger.Info("Receive AddRequest")

	sum := request.A + request.B
	if sum > 100 {
		return ADD_RESULT_TOO_BIG
	}

	s.Response(&proto.AddResponse{Sum: sum})
	return scyna.OK
}
