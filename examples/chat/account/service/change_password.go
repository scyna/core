package account

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/chat/account/domain"
	"github.com/scyna/core/examples/chat/account/proto"
)

func ChangePasswordHandler(cmd *scyna.Command, request *proto.ChangePasswordRequest) scyna.Error {
	cmd.Logger.Info("Receive ChangePaswordRequest")

	repository := domain.LoadRepository(cmd.Logger)
	account, ret := repository.GetAccountByID(request.Id)
	if ret != nil {
		return domain.USER_NOT_EXISTED
	}

	if ret = repository.LoadPassword(account); ret != nil {
		return ret
	}

	if ret = account.ChangePassword(request.Current, request.Future); ret != nil {
		return ret
	}

	if ret = repository.UpdatePassword(account); ret != nil {
		return ret
	}

	return scyna.OK
}
