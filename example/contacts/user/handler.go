package user

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/example/contacts/proto"
	"log"
	"net/http"
)

func HandlerEventMessage(ctx *scyna.Context, message *proto.User) {
	log.Println("X Event")
}

func HandlerSyncMessage(ctx *scyna.Context, message *proto.User) *http.Request {
	log.Println("X Sync " + message.Email)
	return nil
}
