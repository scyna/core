package services

import (
	scyna "github.com/scyna/core"
	PROTO "github.com/scyna/core/example/proto/generated"
	"github.com/scyna/core/example/shared"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateAccountHandler(ctx *scyna.Endpoint, request *PROTO.CreateAccountRequest) scyna.Error {
	ctx.Info("CreateAccountHandler")

	model, err := shared.Store.CreateModel("a@gmail.com")
	if err != nil {
		return err
	}

	model.Data = &PROTO.AccountModel{
		ID:       int64(scyna.ID.Next()),
		Email:    request.Email,
		Name:     request.Name,
		Password: request.Password,
		Created:  timestamppb.Now(),
	}

	if err := model.CommitAndProject(&PROTO.AccountCreated{
		ID:       model.Data.ID,
		Email:    request.Email,
		Name:     request.Name,
		Password: request.Password,
	}); err != nil {
		return err
	}
	ctx.RaiseEvent(model.Event)

	return ctx.OK(&PROTO.CreateAccountResponse{ID: model.Data.ID})
}
