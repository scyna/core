package model

import scyna "github.com/scyna/core"

const ACCOUNT_TABLE = "chat_account.account"

type Account struct {
	LOG      scyna.Logger
	ID       uint64
	Name     string
	Email    EmailAddress
	Password Password
	Gender   Gender
	Tel      PhoneNumber
}

func (acc *Account) ChangePassword(current string, future string) scyna.Error {
	return nil
}
