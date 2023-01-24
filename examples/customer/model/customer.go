package model

type Name string

type Customer struct {
	Identity Identity
	Name     Name
	Gender   Gender
	Email    EmailAddress
	Phone    PhoneNumber
	Address  StreetAddress
}
