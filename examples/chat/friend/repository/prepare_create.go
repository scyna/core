package repository

import (
	"github.com/scylladb/gocqlx/v2/qb"
	scyna "github.com/scyna/core"
)

func ListFriend(LOG scyna.Logger, uid uint64) (*scyna.Error, []*User) {
	var friends []uint64
	var ret []*User

	if err := qb.Select("ex.has_friend").
		Columns("friend_id").
		Where(qb.Eq("user_id")).
		Limit(1).
		Query(scyna.DB).Bind(uid).SelectRelease(friends); err != nil {
		return scyna.SERVER_ERROR, ret
	}

	if len(friends) == 0 {
		return nil, ret
	}

	ret = make([]*User, len(friends))

	qSelect := qb.Select("ex.user").
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
