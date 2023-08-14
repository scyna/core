package example

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/example/PROTO"
)

var store scyna.EventStore[*PROTO.AccountModel] = scyna.NewEventStore[*PROTO.AccountModel]("account")

func CreateAccountHandler(ctx *scyna.Endpoint, request *PROTO.CreateAccountRequest) scyna.Error {
	ctx.Info("CreateAccountHandler")

	model, err := store.CreateModel("a@gmail.com")
	if err != nil {
		return err
	}

	model.Data = &PROTO.AccountModel{
		ID:       int64(scyna.ID.Next()),
		Email:    request.Email,
		Name:     request.Name,
		Password: request.Password,
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

	return nil
}
