package hello_test

import (
	"os"
	"testing"

	scyna "github.com/scyna/core"
	"github.com/scyna/core/example/hello"
	scyna_test "github.com/scyna/core/testing"
)

func TestMain(m *testing.M) {
	scyna_test.Init()

	scyna.RegisterService(hello.HELLO_URL, hello.Hello)
	scyna.RegisterService(hello.ADD_URL, hello.Add)

	exitVal := m.Run()
	scyna_test.Release()
	os.Exit(exitVal)
}
