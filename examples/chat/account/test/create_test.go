package account_test

import (
	"testing"

	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/chat/account/model"
	"github.com/scyna/core/examples/chat/account/proto"
	scyna_test "github.com/scyna/core/testing"
)

func TestCreateShouldReturnSuccess(t *testing.T) {
	cleanup()
	scyna_test.EndpointTest(model.CREATE_USER_URL).
		WithRequest(&proto.Account{
			Email:    "a@gmail.com",
			Name:     "Nguyen Van A",
			Password: "1234565",
		}).
		ExpectSuccess().Run(t)
}

func TestCreateThenGet(t *testing.T) {
	cleanup()
	var response proto.CreateUserResponse
	scyna_test.EndpointTest(model.CREATE_USER_URL).
		WithRequest(&proto.Account{
			Email:    "a@gmail.com",
			Name:     "Nguyen Van A",
			Password: "1234565",
		}).
		ExpectSuccess().Run(t, &response)

	scyna_test.EndpointTest(model.GET_USER_URL).
		WithRequest(&proto.GetUserByEmailRequest{Email: "a@gmail.com"}).
		ExpectResponse(&proto.Account{
			Id:       response.Id,
			Email:    "a@gmail.com",
			Name:     "Nguyen Van A",
			Password: "1234565",
		}).Run(t)
}

func TestCreateBadEmail(t *testing.T) {
	cleanup()
	scyna_test.EndpointTest(model.CREATE_USER_URL).
		WithRequest(&proto.Account{
			Email:    "a+gmail.com",
			Name:     "Nguyen Van A",
			Password: "1234565",
		}).
		ExpectError(scyna.REQUEST_INVALID).Run(t)
}
