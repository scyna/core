package hello_test

import (
	"testing"

	"github.com/scyna/core/example/hello"
	"github.com/scyna/core/example/hello/proto"
	scyna_test "github.com/scyna/core/testing"
)

func TestAddSuccess(t *testing.T) {
	scyna_test.ServiceTest(hello.ADD_URL).
		WithRequest(&proto.AddRequest{A: 5, B: 73}).
		ExpectResponse(&proto.AddResponse{Sum: 78}).
		Run(t)
}

func TestAddTooBig(t *testing.T) {
	scyna_test.ServiceTest(hello.ADD_URL).
		WithRequest(&proto.AddRequest{A: 50, B: 73}).
		ExpectError(hello.ADD_RESULT_TOO_BIG).
		Run(t)
}
