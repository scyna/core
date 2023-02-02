package service

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/account/domain"
	"github.com/scyna/core/examples/account/model"
	proto "github.com/scyna/core/examples/account/proto/generated"
)

func CreateAccountHandler(ctx *scyna.Endpoint, request *proto.CreateAccountRequest) scyna.Error {
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

	command := scyna.NewCommand(&ctx.Context).
		SetAggregateID(account.ID).
		SetChannel(ACCOUNT_CREATED_CHANNEL).
		SetEvent(&proto.AccountCreated{
			Id:    account.ID,
			Name:  account.Name,
			Email: account.Email.String()})

	if ret = repository.CreateAccount(command, &account); ret != nil {
		return ret
	}

	if ret = command.Commit(); ret != nil {
		return ret
	}

	ctx.Response(&proto.CreateAccountResponse{Id: account.ID})

	return scyna.OK
}
