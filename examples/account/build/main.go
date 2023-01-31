package main

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/account/domain"
	"github.com/scyna/core/examples/account/repository"
	"github.com/scyna/core/examples/account/service"
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

	scyna.InitSingleWriter("ex_account")

	scyna.RegisterEndpoint("/scyna.example/user/create", service.CreateAccountHandler)
	scyna.Start()
}
