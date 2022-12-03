package hello_test

import (
	"testing"

	scyna "github.com/scyna/core"
	"github.com/scyna/core/example/hello"
	"github.com/scyna/core/example/hello/proto"
	scyna_test "github.com/scyna/core/testing"
)

func TestHelloSuccess(t *testing.T) {
	scyna_test.ServiceTest(hello.HELLO_URL).
		WithRequest(&proto.HelloRequest{Name: "Alice"}).
		ExpectResponse(&proto.HelloResponse{Content: "Hello Alice"}).
		Run(t)
}

func TestHelloEmptyName(t *testing.T) {
	scyna_test.ServiceTest(hello.HELLO_URL).
		WithRequest(&proto.HelloRequest{}).
		ExpectError(scyna.REQUEST_INVALID).
		Run(t)
}

func TestHelloLongName(t *testing.T) {
	name := "Very long name will cause request invalid."
	scyna_test.ServiceTest(hello.HELLO_URL).
		WithRequest(&proto.HelloRequest{Name: name}).
		ExpectError(scyna.REQUEST_INVALID).
		Run(t)
}

func TestHelloShortName(t *testing.T) {
	scyna_test.ServiceTest(hello.HELLO_URL).
		WithRequest(&proto.HelloRequest{Name: "A"}).
		ExpectError(scyna.REQUEST_INVALID).
		Run(t)
}
