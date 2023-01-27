package account

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/chat/account/domain"
	"github.com/scyna/core/examples/chat/account/model"
	proto "github.com/scyna/core/examples/chat/account/proto/generated"
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
