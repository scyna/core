package repository

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/account/domain"
)

const ACCOUNT_TABLE = "ex_account.account"

type accountRepository struct {
	LOG scyna.Logger
}

func NewRepository(LOG scyna.Logger) domain.IRepository {
	return &accountRepository{LOG: LOG}
}
