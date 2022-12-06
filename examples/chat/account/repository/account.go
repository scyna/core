package repository

import "github.com/scyna/core/examples/chat/account/proto"

const ACCOUNT_TABLE = "chat_account.account"

type Account struct {
	ID       uint64 `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

func (acc *Account) FromDTO(account *proto.Account) {
	acc.ID = account.Id
	acc.Name = account.Name
	acc.Email = account.Email
	acc.Password = account.Password
}

func (acc *Account) ToDTO() *proto.Account {
	return &proto.Account{
		Id:       acc.ID,
		Name:     acc.Name,
		Email:    acc.Email,
		Password: acc.Password,
	}
}
