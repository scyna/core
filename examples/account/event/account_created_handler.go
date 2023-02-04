package event

import (
	"log"

	scyna "github.com/scyna/core"
	proto "github.com/scyna/core/examples/account/proto/generated"
)

func AccountCreatedHandler(ctx *scyna.Context, event *proto.AccountCreated) {
	ctx.Logger.Info("AccountCreated handler")
	log.Print(event)
}
