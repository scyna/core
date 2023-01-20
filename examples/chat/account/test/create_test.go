package account_test

import (
	"testing"

	"github.com/scyna/core/examples/chat/account/domain"
	proto "github.com/scyna/core/examples/chat/account/proto/generated"
	scyna_test "github.com/scyna/core/testing"
)

func TestCreateShouldReturnSuccess(t *testing.T) {
	cleanup()
	scyna_test.EndpointTest(domain.CREATE_USER_URL).
		WithRequest(&proto.CreateAccountRequest{
			Email:    "a@gmail.com",
			Name:     "Nguyen Van A",
			Password: "1234565",
		}).
		ExpectSuccess().Run(t)
}

// func TestCreateThenGet(t *testing.T) {
// 	cleanup()
// 	var response proto.CreateAccountResponse
// 	scyna_test.EndpointTest(domain.CREATE_USER_URL).
// 		WithRequest(&proto.CreateAccountRequest{
// 			Email:    "a@gmail.com",
// 			Name:     "Nguyen Van A",
// 			Password: "1234565",
// 		}).
// 		ExpectSuccess().Run(t, &response)

// 	scyna_test.EndpointTest(domain.GET_USER_URL).
// 		WithRequest(&proto.GetAccountByEmailRequest{Email: "a@gmail.com"}).
// 		ExpectResponse(&proto.Account{
// 			Id:    response.Id,
// 			Email: "a@gmail.com",
// 			Name:  "Nguyen Van A",
// 		}).Run(t)
// }

// func TestCreateBadEmail(t *testing.T) {
// 	cleanup()
// 	scyna_test.EndpointTest(domain.CREATE_USER_URL).
// 		WithRequest(&proto.Account{
// 			Email: "a+gmail.com",
// 			Name:  "Nguyen Van A",
// 		}).
// 		ExpectError(scyna.REQUEST_INVALID).Run(t)
// }
