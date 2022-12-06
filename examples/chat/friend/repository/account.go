package repository

import "github.com/scyna/core/examples/chat/friend/proto"

type Account struct {
	ID    uint64 `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

func (u *Account) ToDTO() *proto.Account {
	return &proto.Account{
		Id:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
