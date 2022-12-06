package account

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/chat/account/model"
	"github.com/scyna/core/examples/chat/account/proto"
	"github.com/scyna/core/examples/chat/account/repository"
)

func CreateAccountHandler(cmd *scyna.Command, request *proto.Account) {
	cmd.Logger.Info("Receive CreateUserRequest")
	if err := validateCreateUserRequest(request); err != nil {
		cmd.Error(scyna.REQUEST_INVALID)
		return
	}

	if err, _ := repository.GetByEmail(cmd.Logger, request.Email); err == nil {
		cmd.Error(model.USER_EXISTED)
		return

	}

	var account repository.Account
	account.FromDTO(request)
	account.ID = scyna.ID.Next()

	repository.PrepareCreate(cmd, &account)

	cmd.Done(&proto.CreateUserResponse{Id: account.ID},
		account.ID,
		model.ACCOUNT_CREATED_CHANNEL,
		&proto.UserCreated{
			Id:    account.ID,
			Name:  account.Name,
			Email: account.Email})
}

func validateCreateUserRequest(user *proto.Account) error {
	return validation.ValidateStruct(user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(5, 10)),
		validation.Field(&user.Name, validation.Required, validation.Length(1, 100)))
}
