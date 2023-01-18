package domain

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/chat/account/model"
)

func CheckAccountExists(repository IRepository, email model.EmailAddress) scyna.Error {
	if _, err := repository.GetAccountByEmail(email); err == nil {
		return USER_EXISTED
	}
	return nil
}
