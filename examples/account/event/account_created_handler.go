package event

import (
	"log"
	"time"

	scyna "github.com/scyna/core"
	proto "github.com/scyna/core/examples/account/proto/generated"
	"github.com/scyna/core/examples/account/service"
)

func AccountCreatedHandler(ctx *scyna.Event, event *proto.AccountCreated) {
	ctx.Logger.Info("AccountCreated handler")
	log.Print(event)
	ctx.ScheduleTask(service.SEND_EMAIL_CHANNEL, time.Now(), 61, &proto.SendEmailTask{
		Email:   event.Email,
		Content: "Just say hello",
	}, 2)
}
