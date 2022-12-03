package main

import scyna "github.com/scyna/core"

const MODULE_CODE = "nats"

func main() {
	scyna.RemoteInit(scyna.RemoteConfig{
		ManagerUrl: "https://localhost:8081",
		Name:       MODULE_CODE,
		Secret:     "123456789aA@#",
	})
	defer scyna.Release()

	scyna.RegisterCommand("/scyna.example/user/create", user.CreateUserHandler)
	scyna.Start()
}
