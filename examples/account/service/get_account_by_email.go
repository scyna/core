package service

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/account/domain"
	"github.com/scyna/core/examples/account/model"
	proto "github.com/scyna/core/examples/account/proto/generated"
)

func GetAccountByEmail(ctx *scyna.Endpoint, request *proto.GetAccountByEmailRequest) scyna.Error {
	ctx.Logger.Info("Receive GetUserRequest")

	repository := domain.LoadRepository(ctx.Logger)

	email, ret := model.ParseEmail(request.Email)
	if ret != nil {
		return ret
	}

	account, ret := repository.GetAccountByEmail(email)
	if ret != nil {
		return ret
	}

	ctx.Response(&proto.Account{
		Id:    account.ID,
		Email: account.Email.String(),
		Name:  account.Name,
		/*TODO*/
	})

	return scyna.OK
}
