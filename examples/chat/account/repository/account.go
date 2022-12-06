package repository

import "github.com/scyna/core/examples/chat/account/proto"

const ACCOUNT_TABLE = "chat_account.account"

type Account struct {
	ID       uint64 `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

func (u *Account) FromDTO(user *proto.Account) {
	u.ID = user.Id
	u.Name = user.Name
	u.Email = user.Email
	u.Password = user.Password
}

func (u *Account) ToDTO() *proto.Account {
	return &proto.Account{
		Id:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}
