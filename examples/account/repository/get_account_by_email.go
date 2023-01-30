package repository

import (
	"log"

	"github.com/scylladb/gocqlx/v2/qb"
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/account/model"
)

func (r *accountRepository) GetAccountByEmail(email model.EmailAddress) (*model.Account, scyna.Error) {

	var account struct {
		ID    uint64 `db:"id"`
		Email string `db:"email"`
		Name  string `db:"name"`
	}

	var tmp = email.String()

	log.Print(tmp)

	if err := qb.Select(ACCOUNT_TABLE).
		Columns("id", "name", "email").
		Where(qb.Eq("email")).
		Limit(1).
		Query(scyna.DB).Bind(tmp).GetRelease(&account); err != nil {
		r.LOG.Error(err.Error())
		return nil, model.USER_NOT_EXISTED
	}

	ret := &model.Account{
		LOG:  r.LOG,
		ID:   account.ID,
		Name: account.Name,
	}

	var err scyna.Error
	if ret.Email, err = model.ParseEmail(account.Email); err != nil {
		return nil, model.BAD_EMAIL
	}

	return ret, nil
}
