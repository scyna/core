package repository

import (
	"github.com/scylladb/gocqlx/v2/qb"
	scyna "github.com/scyna/core"
)

func Create(LOG scyna.Logger, user *User) *scyna.Error {
	if err := qb.Insert(ACCOUNT_TABLE_NAME).
		Columns("id", "name", "email", "password").
		Query(scyna.DB).
		BindStruct(user).
		ExecRelease(); err != nil {
		LOG.Error(err.Error())
		return scyna.SERVER_ERROR
	}
	return nil
}
