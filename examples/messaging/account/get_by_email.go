package account

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/example/messaging/proto"
)

func GetUserByEmail(s *scyna.Endpoint, request *proto.GetUserByEmailRequest) {
	s.Logger.Info("Receive GetUserRequest")
	if err, user := Repository.GetByEmail(s.Logger, request.Email); err != nil {
		s.Error(err)
		return
	} else {
		s.Done(user.ToDTO())
	}
}
