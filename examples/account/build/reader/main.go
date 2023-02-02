package main

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/account/domain"
	"github.com/scyna/core/examples/account/repository"
	"github.com/scyna/core/examples/account/service"
)

func main() {
	scyna.RemoteInit(scyna.RemoteConfig{
		ManagerUrl: "http://localhost:8081",
		Name:       "ex_account",
		Secret:     "12345678",
	})
	defer scyna.Release()

	domain.AttachRepositoryCreator(repository.NewRepository)
	scyna.RegisterEndpoint(service.GET_ACCOUNT_URL, service.GetAccountByEmailHandler)
	scyna.Start()
}
