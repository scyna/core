package account

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/chat/account/domain"
	"github.com/scyna/core/examples/chat/account/model"
	proto "github.com/scyna/core/examples/chat/account/proto/generated"
)

func CreateAccountHandler(ctx *scyna.Command, request *proto.CreateAccountRequest) scyna.Error {
	ctx.Logger.Info("Receive CreateUserRequest")

	repository := domain.LoadRepository(ctx.Logger)

	email, ret := model.ParseEmail(request.Email)
	if ret != nil {
		return ret
	}

	if ret = domain.AssureAccountNotExists(repository, email); ret != nil {
		return ret
	}

	account := model.Account{
		LOG:   ctx.Logger,
		ID:    scyna.ID.Next(),
		Email: email,
		Name:  request.Name, /*TODO: check name*/
	}

	if account.Password, ret = model.ParsePassword(request.Password); ret != nil {
		ctx.Logger.Error("wrong password")
		return ret
	}

	if ret = repository.CreateAccount(ctx, &account); ret != nil {
		return ret
	}

	if ret = ctx.StoreEvent(account.ID, ACCOUNT_CREATED_CHANNEL,
		&proto.AccountCreated{
			Id:    account.ID,
			Name:  account.Name,
			Email: account.Email.ToString()}); ret != nil {
		ctx.Logger.Error("Not OK")
		return ret
	}

	ctx.Response(&proto.CreateAccountResponse{Id: account.ID})

	return scyna.OK
}
