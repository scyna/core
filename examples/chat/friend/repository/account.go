package repository

import "github.com/scyna/core/examples/chat/friend/proto"

const ACCOUNT_TABLE = "chat_friend.account"
const FRIEND_TABLE = "chat_friend.has_friend"

type Account struct {
	ID    uint64 `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

func (acc *Account) ToDTO() *proto.Account {
	return &proto.Account{
		Id:    acc.ID,
		Name:  acc.Name,
		Email: acc.Email,
	}
}
