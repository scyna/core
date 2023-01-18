package repository

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/chat/account/model"
)

func (r *accountRepository) LoadPassword(account *model.Account) scyna.Error {
	return nil
}

func (r *accountRepository) UpdatePassword(cmd *scyna.Command, account *model.Account) scyna.Error {
	return nil
}
