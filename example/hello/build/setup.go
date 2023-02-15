package main

import (
	"ex/hello/service"

	scyna_setup "github.com/scyna/core/setup"
)

func main() {
	scyna_setup.Init()
	scyna_setup.NewModule("ex_hello", "123456").Build()

	scyna_setup.NewClient("hello_test", "123456").
		UseEndpoint(service.ADD_URL).
		UseEndpoint(service.HELLO_URL).
		Build()

}
