package account

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/messaging/account/model"
	"github.com/scyna/core/examples/messaging/account/proto"
)

const CreateUserUrl = "/scyna.example/user/create"

func CreateUserHandler(cmd *scyna.Command, request *proto.User) {
	cmd.Logger.Info("Receive CreateUserRequest")
	if err := validateCreateUserRequest(request); err != nil {
		cmd.Error(scyna.REQUEST_INVALID)
		return
	}

	if err, _ := model.Repository.GetByEmail(cmd.Logger, request.Email); err == nil {
		cmd.Error(model.USER_EXISTED)
		return

	}

	var user model.User
	user.FromDTO(request)
	user.ID = scyna.ID.Next()

	model.Repository.PrepareCreate(cmd, &user)

	cmd.Done(&proto.CreateUserResponse{Id: user.ID},
		user.ID,
		"ex.user.user_created",
		&proto.UserCreated{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email})
}

func validateCreateUserRequest(user *proto.User) error {
	return validation.ValidateStruct(user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(5, 10)),
		validation.Field(&user.Name, validation.Required, validation.Length(1, 100)))
}
