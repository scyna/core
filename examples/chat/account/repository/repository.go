package repository

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/chat/account/domain"
)

const ACCOUNT_TABLE = "chat_account.account"

type accountRepository struct {
	LOG scyna.Logger
}

func LoadAccountRepository(LOG scyna.Logger) domain.IRepository {
	return &accountRepository{LOG: LOG}
}
