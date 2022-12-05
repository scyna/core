package user_test

import (
	"os"
	"testing"

	scyna "github.com/scyna/core"
	"github.com/scyna/core/example/contacts/user"
	scyna_test "github.com/scyna/core/testing"
)

func TestMain(m *testing.M) {
	scyna_test.Init()
	scyna.InitEventStore("ex")

	/*register services*/
	scyna.RegisterCommand(user.CREATE_USER_URL, user.CreateUserHandler)
	scyna.RegisterEndpoint(user.GET_USER_URL, user.GetUserByEmail)

	exitVal := m.Run()
	cleanup()
	scyna_test.Release()
	os.Exit(exitVal)
}

func cleanup() {
	scyna.DB.Query("TRUNCATE ex.user", nil).ExecRelease()
}
