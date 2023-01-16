package model

const ACCOUNT_TABLE = "chat_account.account"

type Account struct {
	ID       uint64
	Name     string
	Email    EmailAddress
	Password Password
	Gender   Gender
	Tel      PhoneNumber
}
