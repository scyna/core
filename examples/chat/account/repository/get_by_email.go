package repository

import (
	"github.com/scylladb/gocqlx/v2/qb"
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/chat/account/model"
)

func GetByEmail(LOG scyna.Logger, email string) (*scyna.Error, *User) {
	var user User
	if err := qb.Select(ACCOUNT_TABLE_NAME).
		Columns("id", "name", "email", "password").
		Where(qb.Eq("email")).
		Limit(1).
		Query(scyna.DB).Bind(email).Get(&user); err != nil {
		LOG.Error(err.Error())
		return model.USER_NOT_EXISTED, nil
	}
	return nil, &user
}
