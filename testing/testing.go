package scyna_test

import (
	"log"

	scyna "github.com/scyna/core"
)

func Init(module string) {
	scyna.RemoteInit(scyna.RemoteConfig{
		ManagerUrl: "http://127.0.0.1:8081",
		Name:       module,
		Secret:     "123456",
	})
	log.Print(scyna.Session.ID())
	scyna.UseDirectLog(1)
}

func Release() {
	scyna.Release()
}
