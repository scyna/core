package service

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/account/domain"
	proto "github.com/scyna/core/examples/account/proto/generated"
)

func ChangePasswordHandler(ctx *scyna.Command, request *proto.ChangePasswordRequest) scyna.Error {
	ctx.Logger.Info("Receive ChangePasswordRequest")

	repository := domain.LoadRepository(ctx.Logger)
	account, ret := repository.GetAccountByID(request.Id)
	if ret != nil {
		return ret
	}

	if ret = repository.LoadPassword(account); ret != nil {
		return ret
	}

	if ret = account.ChangePassword(request.Current, request.Future); ret != nil {
		return ret
	}

	if ret = repository.UpdatePassword(ctx, account); ret != nil {
		return ret
	}

	if ret = ctx.StoreEvent(account.ID, PASSWORD_CHANGED_CHANNEL,
		&proto.PasswordChanged{
			Id:      account.ID,
			Current: request.Current,
			Future:  request.Future}); ret != nil {
		return ret
	}

	return scyna.OK
}
