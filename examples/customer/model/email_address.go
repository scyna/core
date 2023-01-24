package model

type EmailAddress struct {
	value string
}

func (e EmailAddress) String() string {
	return e.value
}

func (e EmailAddress) NotProvided() bool {
	return len(e.value) == 0
}

func NewEmailAddress(address string) EmailAddress {
	/*TODO validate*/
	return EmailAddress{value: address}
}
