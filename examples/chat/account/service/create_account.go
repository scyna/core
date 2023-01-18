package account

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/chat/account/domain"
	"github.com/scyna/core/examples/chat/account/model"
	"github.com/scyna/core/examples/chat/account/proto"
)

func CreateAccountHandler(cmd *scyna.Command, request *proto.CreateAccountRequest) scyna.Error {
	cmd.Logger.Info("Receive CreateUserRequest")

	repository := domain.LoadRepository(cmd.Logger)

	email, ret := model.ParseEmail(request.Email)
	if ret != nil {
		return ret
	}

	if ret = domain.CheckAccountExists(repository, email); ret != nil {
		return ret
	}

	account := model.Account{
		LOG:   cmd.Logger,
		ID:    scyna.ID.Next(),
		Email: email,
		Name:  request.Name, /*TODO: check name*/
	}

	if account.Password, ret = model.ParsePassword(request.Password); ret != nil {
		return ret
	}

	repository.CreateAccount(cmd, &account)

	if ret = cmd.Commit(account.ID, domain.ACCOUNT_CREATED_CHANNEL,
		&proto.UserCreated{
			Id:    account.ID,
			Name:  account.Name,
			Email: account.Email.ToString()}); ret != nil {
		return ret
	}

	cmd.Response(&proto.CreateAccountResponse{Id: account.ID})

	return scyna.OK
}
