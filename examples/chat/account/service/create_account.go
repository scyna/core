package account

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/chat/account/domain"
	"github.com/scyna/core/examples/chat/account/model"
	"github.com/scyna/core/examples/chat/account/proto"
)

func CreateAccountHandler(ctx *scyna.Command, request *proto.CreateAccountRequest) scyna.Error {
	ctx.Logger.Info("Receive CreateUserRequest")

	repository := domain.LoadRepository(ctx.Logger)

	email, ret := model.ParseEmail(request.Email)
	if ret != nil {
		return ret
	}

	if ret = domain.CheckAccountExists(repository, email); ret != nil {
		return ret
	}

	account := model.Account{
		LOG:   ctx.Logger,
		ID:    scyna.ID.Next(),
		Email: email,
		Name:  request.Name, /*TODO: check name*/
	}

	if account.Password, ret = model.ParsePassword(request.Password); ret != nil {
		return ret
	}

	repository.CreateAccount(ctx, &account)

	if ret = ctx.StoreEvent(account.ID, domain.ACCOUNT_CREATED_CHANNEL,
		&proto.AccountCreated{
			Id:    account.ID,
			Name:  account.Name,
			Email: account.Email.ToString()}); ret != nil {
		return ret
	}

	ctx.Response(&proto.CreateAccountResponse{Id: account.ID})

	return scyna.OK
}
