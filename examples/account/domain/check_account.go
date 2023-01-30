package domain

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/account/model"
)

func AssureAccountNotExists(repository IRepository, email model.EmailAddress) scyna.Error {
	if _, err := repository.GetAccountByEmail(email); err == nil {
		return model.USER_EXISTED
	}
	return nil
}
