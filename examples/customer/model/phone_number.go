package model

type PhoneNumber struct {
	value string
}

func (p PhoneNumber) Number() string {
	return p.value
}

func NewPhoneNumber(number string) PhoneNumber {
	/*TODO: validate*/
	return PhoneNumber{value: number}
}
