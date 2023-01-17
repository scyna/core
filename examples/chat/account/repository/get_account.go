package repository

import (
	"github.com/scylladb/gocqlx/v2/qb"
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/chat/account/domain"
	"github.com/scyna/core/examples/chat/account/model"
)

func (r *accountRepository) GetAccount(email model.EmailAddress) (*model.Account, scyna.Error) {
	var user model.Account
	if err := qb.Select(ACCOUNT_TABLE).
		Columns("id", "name", "email", "password").
		Where(qb.Eq("email")).
		Limit(1).
		Query(scyna.DB).Bind(email).GetRelease(&user); err != nil {
		r.LOG.Error(err.Error())
		return nil, domain.USER_NOT_EXISTED
	}
	return &user, nil
}
