package user

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/example/contacts/proto"
)

func GetUserByEmail(s *scyna.Service, request *proto.GetUserByEmailRequest) {
	s.Logger.Info("Receive GetUserRequest")
	if err, user := Repository.GetByEmail(s.Logger, request.Email); err != nil {
		s.Error(err)
		return
	} else {
		s.Done(user.ToDTO())
	}
}
