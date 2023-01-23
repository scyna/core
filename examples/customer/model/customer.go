package model

type Customer struct {
	Identity Identity
	Name     Name
	Email    EmailAddress
	Phone    PhoneNumber
	Address  StreetAddress
}
