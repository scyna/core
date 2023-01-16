package repository

import (
	"github.com/scylladb/gocqlx/v2/qb"
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/chat/account/model"
)

func GetByEmail(LOG scyna.Logger, email string) (scyna.Error, *Account) {
	var user Account
	if err := qb.Select(ACCOUNT_TABLE).
		Columns("id", "name", "email", "password").
		Where(qb.Eq("email")).
		Limit(1).
		Query(scyna.DB).Bind(email).GetRelease(&user); err != nil {
		LOG.Error(err.Error())
		return model.USER_NOT_EXISTED, nil
	}
	return nil, &user
}
