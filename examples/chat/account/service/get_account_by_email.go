package account

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/chat/account/proto"
	"github.com/scyna/core/examples/chat/account/repository"
)

func GetAccountByEmail(s *scyna.Endpoint, request *proto.GetUserByEmailRequest) scyna.Error {
	s.Logger.Info("Receive GetUserRequest")
	if err, user := repository.GetByEmail(s.Logger, request.Email); err != nil {
		return err
	} else {
		s.Done(user.ToDTO())
		return scyna.OK
	}
}
