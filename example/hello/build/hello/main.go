package main

import (
	"ex/hello/service"

	scyna "github.com/scyna/core"
)

const MODULE_CODE = "scyna_test"

func main() {
	scyna.RemoteInit(scyna.RemoteConfig{
		ManagerUrl: "http://localhost:8081",
		Name:       MODULE_CODE,
		Secret:     "123456",
	})
	defer scyna.Release()

	scyna.RegisterEndpoint(service.ADD_URL, service.AddHandler)
	scyna.RegisterEndpoint(service.HELLO_URL, service.HelloHandler)

	scyna.Start()
}
