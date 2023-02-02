package account_test

import (
	"testing"

	"github.com/scyna/core/examples/account/model"
	proto "github.com/scyna/core/examples/account/proto/generated"
	"github.com/scyna/core/examples/account/service"
	scyna_test "github.com/scyna/core/testing"
)

func TestCreateShouldReturnSuccess(t *testing.T) {
	cleanup()
	scyna_test.EndpointTest(service.CREATE_ACCOUNT_URL).
		WithRequest(&proto.CreateAccountRequest{
			Email:    "a@gmail.com",
			Name:     "Nguyen Van A",
			Password: "1234565",
		}).
		PublishEventTo(service.ACCOUNT_CREATED_CHANNEL).
		MatchEvent(&proto.AccountCreated{
			Email: "a@gmail.com",
			Name:  "Nguyen Van A",
		}).
		ExpectSuccess().Run(t)
}

func TestCreateThenGet(t *testing.T) {
	cleanup()
	var response proto.CreateAccountResponse
	scyna_test.EndpointTest(service.CREATE_ACCOUNT_URL).
		WithRequest(&proto.CreateAccountRequest{
			Email:    "a@gmail.com",
			Name:     "Nguyen Van A",
			Password: "1234565",
		}).
		PublishEventTo(service.ACCOUNT_CREATED_CHANNEL).
		ExpectSuccess().Run(t, &response)

	scyna_test.EndpointTest(service.GET_ACCOUNT_URL).
		WithRequest(&proto.GetAccountByEmailRequest{Email: "a@gmail.com"}).
		ExpectResponse(&proto.Account{
			Id:    response.Id,
			Email: "a@gmail.com",
			Name:  "Nguyen Van A",
		}).Run(t)
}

func TestCreateBadEmail(t *testing.T) {
	cleanup()
	scyna_test.EndpointTest(service.CREATE_ACCOUNT_URL).
		WithRequest(&proto.Account{
			Email: "a+gmail.com",
			Name:  "Nguyen Van A",
		}).
		ExpectError(model.BAD_EMAIL).Run(t)
}
