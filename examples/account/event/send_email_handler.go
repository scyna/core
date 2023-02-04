package event

import (
	"log"

	scyna "github.com/scyna/core"
	proto "github.com/scyna/core/examples/account/proto/generated"
)

func SendEmailHandler(ctx *scyna.Context, event *proto.SendEmailTask) {
	ctx.Logger.Info("Receive SendEmailTask")
	log.Print(event)
}
