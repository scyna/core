package account_test

import (
	"os"
	"testing"

	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/chat/account/model"
	account "github.com/scyna/core/examples/chat/account/service"
	scyna_test "github.com/scyna/core/testing"
)

func TestMain(m *testing.M) {
	scyna_test.Init()
	scyna.InitEventStore("chat")

	/*register services*/
	scyna.RegisterCommand(model.CREATE_USER_URL, account.CreateAccountHandler)
	scyna.RegisterEndpoint(model.GET_USER_URL, account.GetAccountByEmail)

	exitVal := m.Run()
	cleanup()
	scyna_test.Release()
	os.Exit(exitVal)
}

func cleanup() {
	scyna.DB.Query("TRUNCATE chat_account.account", nil).ExecRelease()
}
