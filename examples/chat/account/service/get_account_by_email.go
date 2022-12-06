package account

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/chat/account/proto"
	"github.com/scyna/core/examples/chat/account/repository"
)

func GetUserByEmail(s *scyna.Endpoint, request *proto.GetUserByEmailRequest) {
	s.Logger.Info("Receive GetUserRequest")
	if err, user := repository.GetByEmail(s.Logger, request.Email); err != nil {
		s.Error(err)
		return
	} else {
		s.Done(user.ToDTO())
	}
}
