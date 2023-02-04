package main

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/account/event"
	"github.com/scyna/core/examples/account/service"
)

func main() {
	scyna.RemoteInit(scyna.RemoteConfig{
		ManagerUrl: "http://localhost:8081",
		Name:       "ex_account",
		Secret:     "123456",
	})
	defer scyna.Release()

	scyna.RegisterEvent("ex_account", service.ACCOUNT_CREATED_CHANNEL, event.AccountCreatedHandler)
	scyna.RegisterTask("ex_account", service.SEND_EMAIL_CHANNEL, event.SendEmailHandler)
	scyna.Start()
}
