package main

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/chat/account/domain"
	"github.com/scyna/core/examples/chat/account/repository"
	account "github.com/scyna/core/examples/chat/account/service"
)

const MODULE_CODE = "chat_account"

func main() {
	scyna.RemoteInit(scyna.RemoteConfig{
		ManagerUrl: "http://localhost:8081",
		Name:       MODULE_CODE,
		Secret:     "123456789aA@#",
	})
	defer scyna.Release()

	domain.AttachRepositoryCreator(repository.NewRepository)

	scyna.InitSingleWriter("chat_account")

	scyna.RegisterCommand("/scyna.example/user/create", account.CreateAccountHandler)
	scyna.Start()
}
