package account

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/chat/account/domain"
	"github.com/scyna/core/examples/chat/account/model"
	"github.com/scyna/core/examples/chat/account/proto"
)

func GetAccountByEmail(s *scyna.Endpoint, request *proto.GetUserByEmailRequest) scyna.Error {
	s.Logger.Info("Receive GetUserRequest")

	repository := domain.LoadAccountRepository(s.Logger)

	email, ret := model.ParseEmail(request.Email)
	if ret != nil {
		return scyna.REQUEST_INVALID
	}

	account, ret := repository.GetAccount(email)
	if ret != nil {
		return domain.USER_NOT_EXISTED
	}

	s.Done(&proto.Account{
		Email: account.Email.ToString(),
		Name:  account.Name,
		/*TODO*/
	})

	return scyna.OK
}
