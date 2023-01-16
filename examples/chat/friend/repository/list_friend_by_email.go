package repository

import (
	"github.com/scylladb/gocqlx/v2/qb"
	scyna "github.com/scyna/core"
)

func ListFriendByEmail(LOG scyna.Logger, uid uint64) (scyna.Error, []*Account) {
	var friends []uint64
	var ret []*Account

	if err := qb.Select(FRIEND_TABLE).
		Columns("friend_id").
		Where(qb.Eq("account_id")).
		Limit(1).
		Query(scyna.DB).Bind(uid).SelectRelease(friends); err != nil {
		return scyna.SERVER_ERROR, ret
	}

	if len(friends) == 0 {
		return nil, ret
	}

	ret = make([]*Account, len(friends))

	qSelect := qb.Select(ACCOUNT_TABLE).
		Columns("id", "name", "email").
		Where(qb.Eq("id")).
		Limit(1).
		Query(scyna.DB)

	for i, id := range friends {
		qSelect.Bind(id).Get(ret[i])
	}

	qSelect.Release()
	return nil, ret
}
