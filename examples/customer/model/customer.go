package model

type Name string

type Customer struct {
	Identity Identity
	Name     Name
	Email    EmailAddress
	Phone    PhoneNumber
	Address  StreetAddress
}
