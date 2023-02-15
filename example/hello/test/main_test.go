package hello_test

import (
	hello "ex/hello/service"
	"os"
	"testing"

	scyna "github.com/scyna/core"
	scyna_test "github.com/scyna/core/testing"
)

func TestMain(m *testing.M) {
	scyna_test.Init("scyna_test")

	scyna.RegisterEndpoint(hello.HELLO_URL, hello.HelloHandler)
	scyna.RegisterEndpoint(hello.ADD_URL, hello.AddHandler)

	exitVal := m.Run()
	scyna_test.Release()
	os.Exit(exitVal)
}
