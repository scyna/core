package main

import (
	scyna "github.com/scyna/core"
)

const MODULE_CODE = "scyna.test"

func main() {
	scyna.RemoteInit(scyna.RemoteConfig{
		ManagerUrl: "http://localhost:8081",
		Name:       MODULE_CODE,
		Secret:     "123456",
	})
	defer scyna.Release()

	scyna.Start()
}
