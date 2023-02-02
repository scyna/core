package service

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/account/domain"
	proto "github.com/scyna/core/examples/account/proto/generated"
)

func ChangePasswordHandler(ctx *scyna.Endpoint, request *proto.ChangePasswordRequest) scyna.Error {
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

	command := scyna.NewCommand(&ctx.Context).
		SetAggregateID(account.ID).
		SetChannel(PASSWORD_CHANGED_CHANNEL).
		SetEvent(&proto.PasswordChanged{
			Id:      account.ID,
			Current: request.Current,
			Future:  request.Future})

	if ret = repository.UpdatePassword(command, account); ret != nil {
		return ret
	}

	if ret = command.Commit(); ret != nil {
		return ret
	}

	return scyna.OK
}
