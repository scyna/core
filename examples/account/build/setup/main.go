package main

import scyna_setup "github.com/scyna/core/setup"

func main() {
	scyna_setup.Init()
	scyna_setup.NewModule("ex_account", "123456").
		AddEventChannel("ex_account").
		Build()
}
