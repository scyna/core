package model

type Name string
type CustomerID string

type Customer struct {
	ID       CustomerID
	Identity Identity
	Name     CustomerName
	Gender   Gender
	Email    EmailAddress
	Phone    PhoneNumber
	Address  StreetAddress
	DOB      DateOfBirth
}
