package account

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/messaging/account/model"
	"github.com/scyna/core/examples/messaging/account/proto"
)

func GetUserByEmail(s *scyna.Endpoint, request *proto.GetUserByEmailRequest) {
	s.Logger.Info("Receive GetUserRequest")
	if err, user := model.Repository.GetByEmail(s.Logger, request.Email); err != nil {
		s.Error(err)
		return
	} else {
		s.Done(user.ToDTO())
	}
}
