package account_test

import (
	"os"
	"testing"

	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/chat/account/domain"
	"github.com/scyna/core/examples/chat/account/repository"
	account "github.com/scyna/core/examples/chat/account/service"
	scyna_test "github.com/scyna/core/testing"
)

func TestMain(m *testing.M) {
	scyna_test.Init()
	scyna.InitEventStore("chat_account")

	/*register services*/
	scyna.RegisterCommand(domain.CREATE_USER_URL, account.CreateAccountHandler)
	scyna.RegisterEndpoint(domain.GET_USER_URL, account.GetAccountByEmail)

	domain.AttachRepositoryCreator(repository.NewRepository)

	exitVal := m.Run()
	cleanup()
	scyna_test.Release()
	os.Exit(exitVal)
}

func cleanup() {
	scyna.DB.Query("TRUNCATE chat_account.account", nil).ExecRelease()
}
